<template>
  <router-view />
  <template v-if="state.notificationList.length > 0">
    <BBNotification
      :placement="'BOTTOM_RIGHT'"
      :notificationList="state.notificationList"
      @close="removeNotification"
    />
  </template>
</template>

<script lang="ts">
import { reactive, watchEffect, onErrorCaptured } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { isDev } from "./utils";
import { Notification } from "./types";
import { BBNotificationItem } from "./bbkit/types";

// Show at most 3 notifications to prevent excessive notification when shit hits the fan.
const MAX_NOTIFICATION_DISPLAY_COUNT = 3;

// Check expiration every 30 sec and logout if expired
const CHECK_LOGGEDIN_STATE_DURATION = 30 * 1000;

const NOTIFICAITON_DURATION = 6000;
const CRITICAL_NOTIFICAITON_DURATION = 10000;

interface LocalState {
  notificationList: BBNotificationItem[];
  prevLoggedIn: boolean;
}

export default {
  name: "App",
  components: {},
  setup(props, ctx) {
    const store = useStore();
    const router = useRouter();

    const state = reactive<LocalState>({
      notificationList: [],
      prevLoggedIn: store.getters["auth/isLoggedIn"](),
    });

    setInterval(() => {
      const loggedIn = store.getters["auth/isLoggedIn"]();
      if (state.prevLoggedIn != loggedIn) {
        state.prevLoggedIn = loggedIn;
        if (!loggedIn) {
          store.dispatch("auth/logout").then(() => {
            router.push({ name: "auth.signin" });
          });
        }
      }
    }, CHECK_LOGGEDIN_STATE_DURATION);

    const removeNotification = (item: BBNotificationItem) => {
      const index = state.notificationList.indexOf(item);
      if (index >= 0) {
        state.notificationList.splice(index, 1);
      }
    };

    const watchNotification = () => {
      store
        .dispatch("notification/tryPopNotification", {
          module: "bytebase",
        })
        .then((notification: Notification | undefined) => {
          if (notification) {
            if (
              state.notificationList.length < MAX_NOTIFICATION_DISPLAY_COUNT
            ) {
              const item: BBNotificationItem = {
                style: notification.style,
                title: notification.title,
                description: notification.description || "",
                link: notification.link || "",
                linkTitle: notification.linkTitle || "",
              };
              state.notificationList.push(item);
              if (!notification.manualHide) {
                setTimeout(
                  () => {
                    removeNotification(item);
                  },
                  notification.style == "CRITICAL"
                    ? CRITICAL_NOTIFICAITON_DURATION
                    : NOTIFICAITON_DURATION
                );
              }
            }
          }
        });
    };

    watchEffect(watchNotification);

    onErrorCaptured((e: any, _, info) => {
      // If e has response, then we assume it's an http error and has already been
      // handled by the axios global handler.
      if (!e.response) {
        store.dispatch("notification/pushNotification", {
          module: "bytebase",
          style: "CRITICAL",
          title: `Internal error occured`,
          description: isDev() ? e.stack : undefined,
        });
      }
      return true;
    });

    return {
      state,
      removeNotification,
    };
  },
};
</script>
