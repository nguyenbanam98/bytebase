<template>
  <div class="mx-auto w-full max-w-sm">
    <div>
      <img class="h-12 w-auto" src="../../assets/logo.svg" alt="Bytebase" />
      <h2 class="mt-6 text-3xl leading-9 font-extrabold text-main">
        <template v-if="needAdminSetup" class="text-accent font-semnibold">
          Setup
          <span class="text-accent font-semnibold">admin account</span>
        </template>
        <template v-else> Register your account </template>
      </h2>
    </div>

    <div class="mt-8">
      <div class="mt-6">
        <form @submit.prevent="trySignup" class="space-y-6">
          <div>
            <label
              for="email"
              class="block text-sm font-medium leading-5 text-control"
            >
              Email <span class="text-red-600">*</span>
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <input
                id="email"
                type="email"
                v-model="state.email"
                required
                placeholder="jim@example.com"
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
                @input="onTextEmail"
              />
            </div>
          </div>

          <div>
            <label
              for="password"
              class="block text-sm font-medium leading-5 text-control"
            >
              Password <span class="text-red-600">*</span>
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <input
                id="password"
                type="password"
                autocomplete="off"
                :value="state.password"
                @input="changePassword($event.target.value)"
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
              for="password-confirm"
              class="block text-sm font-medium leading-5 text-control"
            >
              Confirm Password
              <span class="text-red-600"
                >* {{ state.showPasswordMismatchError ? "mismatch" : "" }}</span
              >
            </label>
            <div class="mt-1 rounded-md shadow-sm">
              <input
                id="password-confirm"
                type="password"
                autocomplete="off"
                :placeholder="'Confirm password'"
                :value="state.passwordConfirm"
                @input="changePasswordConfirm($event.target.value)"
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
                @input="onTextName"
              />
            </div>
          </div>

          <div>
            <span class="block w-full rounded-md shadow-sm">
              <button
                type="submit"
                :disabled="!allowSignup"
                class="btn-primary w-full flex justify-center py-2 px-4"
              >
                {{ needAdminSetup ? "Create admin account" : "Register" }}
              </button>
            </span>
          </div>
        </form>
      </div>
    </div>

    <div v-if="!needAdminSetup" class="mt-6 relative">
      <div class="absolute inset-0 flex items-center" aria-hidden="true">
        <div class="w-full border-t border-control-border"></div>
      </div>
      <div class="relative flex justify-center text-sm">
        <span class="pl-2 bg-white text-control">
          Already have an account?
        </span>
        <router-link to="/auth/signin" class="accent-link px-2 bg-white">
          Sign in
        </router-link>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { computed, onUnmounted, reactive } from "vue";
import { useStore } from "vuex";
import { useRouter } from "vue-router";
import { SignupInfo, TEXT_VALIDATION_DELAY } from "../../types";
import { isValidEmail } from "../../utils";

interface LocalState {
  email: string;
  password: string;
  passwordConfirm: string;
  passwordValidationTimer?: ReturnType<typeof setTimeout>;
  showPasswordMismatchError: boolean;
  name: string;
  nameManuallyEdited: boolean;
}

export default {
  name: "Signup",
  setup(props, ctx) {
    const store = useStore();
    const router = useRouter();

    const state = reactive<LocalState>({
      email: "",
      password: "",
      passwordConfirm: "",
      showPasswordMismatchError: false,
      name: "",
      nameManuallyEdited: false,
    });

    onUnmounted(() => {
      if (state.passwordValidationTimer) {
        clearInterval(state.passwordValidationTimer);
      }
    });

    const needAdminSetup = computed(() => {
      return store.getters["actuator/needAdminSetup"]();
    });

    const allowSignup = computed(() => {
      return (
        isValidEmail(state.email) &&
        state.password &&
        !state.showPasswordMismatchError
      );
    });

    const passwordMatch = computed(() => {
      return state.password == state.passwordConfirm;
    });

    const refreshPasswordValidation = () => {
      if (state.passwordValidationTimer) {
        clearInterval(state.passwordValidationTimer);
      }

      if (passwordMatch.value) {
        state.showPasswordMismatchError = false;
      } else {
        state.passwordValidationTimer = setTimeout(() => {
          // If error is already displayed, we hide the error only if there is valid input.
          // Otherwise, we hide the error if input is either empty or valid.
          if (state.showPasswordMismatchError) {
            state.showPasswordMismatchError = !passwordMatch.value;
          } else {
            state.showPasswordMismatchError =
              state.password != "" &&
              state.passwordConfirm != "" &&
              !passwordMatch.value;
          }
        }, TEXT_VALIDATION_DELAY);
      }
    };

    const changePassword = (value: string) => {
      state.password = value;
      refreshPasswordValidation();
    };

    const changePasswordConfirm = (value: string) => {
      state.passwordConfirm = value;
      refreshPasswordValidation();
    };

    const onTextEmail = () => {
      const email = state.email.trim();
      if (!state.nameManuallyEdited) {
        const emailParts = email.split("@");
        if (emailParts.length > 0) {
          if (emailParts[0].length > 0) {
            const name = emailParts[0].replace("_", ".");
            const nameParts = name.split(".");
            if (nameParts.length >= 2) {
              state.name = [
                nameParts[0].charAt(0).toUpperCase() + nameParts[0].slice(1),
                nameParts[1].charAt(0).toUpperCase() + nameParts[1].slice(1),
              ].join(" ");
            } else {
              state.name = name.charAt(0).toUpperCase() + name.slice(1);
            }
          }
        }
      }
    };

    const onTextName = () => {
      const name = state.name.trim();
      state.nameManuallyEdited = name.length > 0;
    };

    const trySignup = () => {
      if (!passwordMatch.value) {
        state.showPasswordMismatchError = true;
      } else {
        const signupInfo: SignupInfo = {
          email: state.email,
          password: state.password,
          name: state.name,
        };
        store.dispatch("auth/signup", signupInfo).then(() => {
          router.push("/");
        });
      }
    };

    return {
      state,
      needAdminSetup,
      allowSignup,
      changePassword,
      changePasswordConfirm,
      onTextEmail,
      onTextName,
      trySignup,
    };
  },
};
</script>
