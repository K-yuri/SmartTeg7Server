<template>
  <div class="flex justify-center items-center h-screen bg-gray-200 px-6">
    <div class="p-6 max-w-sm w-full bg-white shadow-md rounded-md">
      <div class="flex justify-center items-center">

        <span class="text-gray-600 font-semibold text-2xl">물품방지태그</span>
      </div>

      <form class="mt-4" @submit.prevent="login">
        <label class="block">
          <span class="text-gray-700 text-sm">ID</span>
          <input
            type="email"
            class="mt-1 block w-full rounded-md focus:border-indigo-600"
            v-model="email"
          />
        </label>

        <label class="block mt-3">
          <span class="text-gray-700 text-sm">Password</span>
          <input
            type="password"
            class="mt-1 block w-full rounded-md focus:border-indigo-600"
            v-model="password"
          />
        </label>

        <div class="flex justify-between items-center mt-4">
          <div>
            <label class="inline-flex items-center">
              <input type="checkbox" class="form-checkbox text-indigo-600" />
              <span class="mx-2 text-gray-600 text-sm">save</span>
            </label>
          </div>

          <div>
            <a
              class="block text-sm fontme text-indigo-700 hover:underline"
              href="#"
              >계정찾기</a
            >
          </div>
        </div>

        <div class="mt-6">
          <button
            type="submit"
            class="py-2 px-4 text-center bg-indigo-600 rounded-md w-full text-white text-sm hover:bg-indigo-500"
          >
           Login
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { defineComponent, ref } from "vue";
import { useRouter } from "vue-router";

export default defineComponent({
  setup() {
    const router = useRouter();
    const email = ref("sss@naver.com");
    const password = ref("12345678");

    function login() {
      router.push("/dashboard");
    }

    return {
      login,
      email,
      password,
    };
  },
});
</script>
