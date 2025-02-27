<template>
  <div class="mx-auto w-full max-w-sm">
    <div>
      <img class="h-12 w-auto" src="../../assets/logo.svg" alt="Bytebase" />
      <h2 class="mt-6 text-3xl leading-9 font-extrabold text-main">
        Activate your
        <span class="text-accent font-semnibold">{{
          state.role.charAt(0).toUpperCase() + state.role.slice(1).toLowerCase()
        }}</span>
        account
      </h2>
    </div>

    <div class="mt-8">
      <div class="mt-6">
        <form @submit.prevent="tryActivate" class="space-y-6">
          <div>
            <label
              for="email"
              class="block text-sm font-medium leading-5 text-control"
            >
              Email
            </label>
            <div class="mt-2 text-base font-medium leading-5 text-accent">
              {{ state.email }}
            </div>
          </div>
          <div>
            <label
              for="password"
              class="block text-sm font-medium leading-5 text-control"
            >
              Password<span class="text-red-600">*</span>
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <input
                id="password"
                type="password"
                autocomplete="on"
                v-model="state.password"
                required
                class="
                  appearance-none
                  block
                  w-full
                  px-3
                  py-2
                  border border-control-border
                  rounded-md
                  placeholder-control-placeholder
                  focus:outline-none
                  focus:shadow-outline-blue
                  focus:border-control-border
                  sm:text-sm sm:leading-5
                "
              />
            </div>
          </div>

          <div>
            <label
              for="email"
              class="block text-sm font-medium leading-5 text-control"
            >
              Name
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <input
                id="name"
                type="text"
                v-model="state.name"
                placeholder="Jim Gray"
                class="
                  appearance-none
                  block
                  w-full
                  px-3
                  py-2
                  border border-control-border
                  rounded-md
                  placeholder-control-placeholder
                  focus:outline-none
                  focus:shadow-outline-blue
                  focus:border-control-border
                  sm:text-sm sm:leading-5
                "
              />
            </div>
          </div>

          <div>
            <span class="block w-full rounded-md shadow-sm">
              <button
                type="submit"
                class="btn-success w-full flex justify-center py-2 px-4"
              >
                Activate
              </button>
            </span>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { reactive } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { ActivateInfo, RoleType } from "../../types";

interface LocalState {
  email: string;
  password: string;
  name: string;
  role: RoleType;
}

export default {
  name: "Activate",
  setup(props, ctx) {
    const store = useStore();
    const router = useRouter();
    const token = router.currentRoute.value.query.token as string;

    // TODO(tianzhou): Get info from activate token
    const state = reactive<LocalState>({
      email: "bob@example.com",
      password: "",
      name: "Bob Invited",
      role: "DEVELOPER",
    });

    const tryActivate = () => {
      const activateInfo: ActivateInfo = {
        email: state.email,
        password: state.password,
        name: state.name,
        token: token,
      };
      store.dispatch("auth/activate", activateInfo).then(() => {
        router.push("/");
      });
    };

    return {
      state,
      tryActivate,
    };
  },
};
</script>
