<template>
  <div class="h-full flex overflow-hidden">
    <!-- Off-canvas menu for mobile, show/hide based on off-canvas menu state. -->
    <div v-if="state.showMobileOverlay" class="md:hidden">
      <div class="fixed inset-0 flex z-40">
        <div class="fixed inset-0">
          <div class="absolute inset-0 bg-gray-600 opacity-75"></div>
        </div>
        <div
          tabindex="0"
          class="
            relative
            flex-1 flex flex-col
            max-w-xs
            w-full
            bg-white
            focus:outline-none
          "
        >
          <div class="absolute top-0 right-0 -mr-12 pt-2">
            <button
              @click.prevent="state.showMobileOverlay = false"
              type="button"
              class="
                ml-1
                flex
                items-center
                justify-center
                h-10
                w-10
                rounded-full
                focus:outline-none
                focus:ring-2
                focus:ring-inset
                focus:ring-white
              "
            >
              <span class="sr-only">Close sidebar</span>
              <!-- Heroicon name: x -->
              <svg
                class="h-6 w-6 text-white"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                aria-hidden="true"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M6 18L18 6M6 6l12 12"
                />
              </svg>
            </button>
          </div>
          <!-- Mobile Sidebar -->
          <div class="flex-1 h-0 py-4 overflow-y-auto">
            <router-view name="leftSidebar" />
          </div>
          <router-link
            to="/archive"
            class="outline-item group flex items-center px-4 pt-4 pb-2"
          >
            <svg
              class="w-5 h-5 mr-2"
              fill="currentColor"
              viewBox="0 0 20 20"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path d="M4 3a2 2 0 100 4h12a2 2 0 100-4H4z"></path>
              <path
                fill-rule="evenodd"
                d="M3 8h14v7a2 2 0 01-2 2H5a2 2 0 01-2-2V8zm5 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z"
                clip-rule="evenodd"
              ></path>
            </svg>
            Archive
          </router-link>
          <div
            class="flex-shrink-0 flex border-t border-block-border px-4 py-2"
          >
            <a
              href="https://github.com/bytebase/bytebase/discussions"
              target="_blank"
              class="
                flex
                justify-between
                items-center
                w-full
                flex-shrink-0
                text-main
                group
              "
            >
              <div class="flex items-center">
                <svg
                  class="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                  ></path>
                </svg>
                <span class="ml-1 text-sm">Help</span>
              </div>
              <div class="text-sm">{{ version }}</div>
            </a>
          </div>
        </div>
        <div class="flex-shrink-0 w-14" aria-hidden="true">
          <!-- Force sidebar to shrink to fit close icon -->
        </div>
      </div>
    </div>

    <!-- Static sidebar for desktop -->
    <aside class="hidden md:flex md:flex-shrink-0">
      <div class="flex flex-col w-52">
        <!-- Sidebar component, swap this element with another sidebar if you like -->
        <div class="flex-1 flex flex-col py-2 overflow-y-auto">
          <router-view name="leftSidebar" />
        </div>
        <router-link
          to="/archive"
          class="outline-item group flex items-center px-4 pt-4 pb-2"
        >
          <svg
            class="w-5 h-5 mr-2"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path d="M4 3a2 2 0 100 4h12a2 2 0 100-4H4z"></path>
            <path
              fill-rule="evenodd"
              d="M3 8h14v7a2 2 0 01-2 2H5a2 2 0 01-2-2V8zm5 3a1 1 0 011-1h2a1 1 0 110 2H9a1 1 0 01-1-1z"
              clip-rule="evenodd"
            ></path>
          </svg>
          Archive
        </router-link>
        <div
          v-if="showQuickstart"
          class="
            flex-shrink-0 flex
            justify-center
            border-t border-block-border
            py-2
          "
        >
          <Quickstart />
        </div>
        <div class="flex-shrink-0 flex border-t border-block-border px-4 py-2">
          <a
            href="https://github.com/bytebase/bytebase/discussions"
            target="_blank"
            class="
              flex
              justify-between
              items-center
              flex-shrink-0
              w-full
              text-main
              group
            "
          >
            <div class="flex items-center py-1">
              <svg
                class="w-5 h-5"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M8.228 9c.549-1.165 2.03-2 3.772-2 2.21 0 4 1.343 4 3 0 1.4-1.278 2.575-3.006 2.907-.542.104-.994.54-.994 1.093m0 3h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                ></path>
              </svg>
              <span class="ml-1 text-sm">Help</span>
            </div>
            <div class="text-sm">{{ version }}</div>
          </a>
        </div>
      </div>
    </aside>
    <div
      class="flex flex-col min-w-0 flex-1 border-l border-r border-block-border"
    >
      <!-- Static sidebar for mobile -->
      <aside class="md:hidden">
        <div
          class="
            flex
            items-center
            justify-start
            bg-gray-50
            border-b border-block-border
            px-4
            py-1.5
          "
        >
          <div>
            <button
              @click.prevent="state.showMobileOverlay = true"
              type="button"
              class="
                -mr-3
                h-12
                w-12
                inline-flex
                items-center
                justify-center
                rounded-md
                text-gray-500
                hover:text-gray-900
              "
            >
              <span class="sr-only">Open sidebar</span>
              <!-- Heroicon name: menu -->
              <svg
                class="h-6 w-6"
                xmlns="http://www.w3.org/2000/svg"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                aria-hidden="true"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M4 6h16M4 12h16M4 18h16"
                />
              </svg>
            </button>
          </div>
          <div v-if="showBreadcrumb" class="ml-4">
            <Breadcrumb />
          </div>
        </div>
      </aside>
      <div class="w-full mx-auto md:flex">
        <div class="md:min-w-0 md:flex-1">
          <div v-if="showBreadcrumb" class="hidden md:block px-4 pt-4">
            <Breadcrumb />
          </div>
          <div v-if="quickActionList.length > 0" class="mx-4 mt-4">
            <QuickActionPanel :quickActionList="quickActionList" />
          </div>
        </div>
      </div>
      <!-- This area may scroll -->
      <div
        class="md:min-w-0 md:flex-1 overflow-y-auto"
        :class="showBreadcrumb || quickActionList.length > 0 ? 'mt-2' : ''"
      >
        <!-- Start main area-->
        <router-view name="content" />
        <!-- End main area -->
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, reactive } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import Breadcrumb from "../components/Breadcrumb.vue";
import Quickstart from "../components/Quickstart.vue";
import QuickActionPanel from "../components/QuickActionPanel.vue";
import { QuickActionType } from "../types";
import { isDBA, isDeveloper, isOwner } from "../utils";

interface LocalState {
  showMobileOverlay: boolean;
}

export default {
  name: "BodyLayout",
  components: {
    Breadcrumb,
    Quickstart,
    QuickActionPanel,
  },
  setup(props, ctx) {
    const store = useStore();
    const router = useRouter();

    const state = reactive<LocalState>({
      showMobileOverlay: false,
    });

    const currentUser = computed(() => store.getters["auth/currentUser"]());

    const quickActionList = computed(() => {
      const role = currentUser.value.role;
      const quickActionListFunc =
        router.currentRoute.value.meta.quickActionListByRole;
      const listByRole = quickActionListFunc
        ? quickActionListFunc(router.currentRoute.value)
        : new Map();
      const list: QuickActionType[] = [];

      // We write this way because for free version, the user wears the three role hat,
      // and we want to display all quick actions relevant to those three roles without duplication.
      if (isOwner(role)) {
        for (const item of listByRole.get("OWNER") || []) {
          list.push(item);
        }
      }

      if (isDBA(role)) {
        for (const item of listByRole.get("DBA") || []) {
          if (
            !list.find((myItem: QuickActionType) => {
              return item == myItem;
            })
          ) {
            list.push(item);
          }
        }
      }

      if (isDeveloper(role)) {
        for (const item of listByRole.get("DEVELOPER") || []) {
          if (
            !list.find((myItem: QuickActionType) => {
              return item == myItem;
            })
          ) {
            list.push(item);
          }
        }
      }
      return list;
    });

    const showBreadcrumb = computed(() => {
      const name = router.currentRoute.value.name;
      return !(name === "workspace.home" || name === "workspace.profile");
    });

    const showQuickstart = computed(() => {
      // Do not show quickstart in demo mode since we don't expect user to alter the data
      return (
        !store.getters["actuator/isDemo"]() &&
        !store.getters["uistate/introStateByKey"]("hidden")
      );
    });

    const version = computed(() => {
      const v = store.getters["actuator/version"]();
      if (v.split(".").length == 3) {
        return `v${v}`;
      }
      return v;
    });

    return {
      state,
      quickActionList,
      showBreadcrumb,
      showQuickstart,
      version,
    };
  },
};
</script>
