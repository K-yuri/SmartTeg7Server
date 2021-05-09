<template>
  <div>
    <h3 class="text-gray-700 text-3xl font-medium">전체 보기</h3>

    <div class="mt-4">
      <div class="flex flex-wrap -mx-6">
        <div class="w-full px-6 sm:w-1/2 xl:w-1/3">
          <div
            class="flex items-center px-5 py-6 shadow-sm rounded-md bg-white"
          >
            <div class="p-3 rounded-full bg-indigo-600 bg-opacity-75">
              <svg
                class="h-8 w-8 text-white"
                viewBox="0 0 28 30"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  
                  fill="currentColor"
                />
                <path
                 
                  fill="currentColor"
                />
                <path
                  
                  fill="currentColor"
                />
                <path
                  
                  fill="currentColor"
                />
                <path
                  
                  fill="currentColor"
                />
                <path
                 
                  fill="currentColor"
                />
              </svg>
            </div>

            <div class="mx-5">
              <h4 class="text-2xl font-semibold text-gray-700">4</h4>
              <div class="text-gray-500">사용중인 회원</div>
            </div>
          </div>
        </div>

        <div class="w-full mt-6 px-6 sm:w-1/2 xl:w-1/3 sm:mt-0">
          <div
            class="flex items-center px-5 py-6 shadow-sm rounded-md bg-white"
          >
            <div class="p-3 rounded-full bg-orange-600 bg-opacity-75">
              <svg
                class="h-8 w-8 text-white"
                viewBox="0 0 28 28"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  
                  fill="currentColor"
                />
                <path
                 
                  fill="currentColor"
                />
                <path
                  
                  fill="currentColor"
                />
              </svg>
            </div>

            <div class="mx-5">
              <h4 class="text-2xl font-semibold text-gray-700">11</h4>
              <div class="text-gray-500">사용중인 태그</div>
            </div>
          </div>
        </div>

        <div class="w-full mt-6 px-6 sm:w-1/2 xl:w-1/3 xl:mt-0">
          <div
            class="flex items-center px-5 py-6 shadow-sm rounded-md bg-white"
          >
            <div class="p-3 rounded-full bg-pink-600 bg-opacity-75">
              <svg
                class="h-8 w-8 text-white"
                viewBox="0 0 28 28"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                 
                  fill="currentColor"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linejoin="round"
                />
                <path
                  
                  stroke="currentColor"
                  stroke-width="2"
                />
              </svg>
            </div>

            <div class="mx-5">
              <h4 class="text-2xl font-semibold text-gray-700">1</h4>
              <div class="text-gray-500">분실 신고</div>
            </div>
          </div>
        </div>
      </div>
    </div>

    <div class="mt-8"></div>

    <div class="flex flex-col mt-8">
      <div class="-my-2 py-2 overflow-x-auto sm:-mx-6 sm:px-6 lg:-mx-8 lg:px-8">
        <div
          class="align-middle inline-block min-w-full shadow overflow-hidden sm:rounded-lg border-b border-gray-200"
        >
          <table class="min-w-full">
            <thead>
              <tr>
                <th
                  class="px-6 py-3 border-b border-gray-200 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider"
                >
                  이름
                </th>
                <th
                  class="px-6 py-3 border-b border-gray-200 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider"
                >
                  태그
                </th>
                <th
                  class="px-6 py-3 border-b border-gray-200 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider"
                >
                  상태
                </th>
                <th
                  class="px-6 py-3 border-b border-gray-200 bg-gray-50 text-left text-xs leading-4 font-medium text-gray-500 uppercase tracking-wider"
                >
                  ---
                </th>
                <th class="px-6 py-3 border-b border-gray-200 bg-gray-50"></th>
              </tr>
            </thead>

            <tbody class="bg-white">
              <tr v-for="(u, index) in users" :key="index">
                <td
                  class="px-6 py-4 whitespace-nowrap border-b border-gray-200"
                >
                  <div class="flex items-center">
                    <div class="flex-shrink-0 h-10 w-10">
                      <img
                        class="h-10 w-10 rounded-full"
                        src="/src/people.jpeg"
                        alt=""
                      />
                    </div>

                    <div class="ml-4">
                      <div class="text-sm leading-5 font-medium text-gray-900">
                        {{ u.name }}
                      </div>
                      <div class="text-sm leading-5 text-gray-500">
                        {{ u.email }}
                      </div>
                    </div>
                  </div>
                </td>

                <td
                  class="px-6 py-4 whitespace-nowrap border-b border-gray-200"
                >
                  <div class="text-sm leading-5 text-gray-900">
                    {{ u.title }}
                  </div>
                  <div class="text-sm leading-5 text-gray-500">
                    {{ u.title2 }}
                  </div>
                </td>

                <td
                  class="px-6 py-4 whitespace-nowrap border-b border-gray-200"
                >
                  <span
                    class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full bg-red-100 text-green-800"
                    >{{ u.status }}</span
                  >
                </td>

                <td
                  class="px-6 py-4 whitespace-nowrap border-b border-gray-200 text-sm leading-5 text-gray-500"
                >
                  {{ u.role }}
                </td>

                <td
                  class="px-6 py-4 whitespace-nowrap text-right border-b border-gray-200 text-sm leading-5 font-medium"
                >
                  <a href="#" class="text-indigo-600 hover:text-indigo-900"
                    >수정</a
                  >
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";

interface User {
  name: string;
  email: string;
  title: string;
  title2: string;
  status: string;
  role: string;
}

export default defineComponent({
  setup() {
    const testUser: User = {
      name: "김아몰랑",
      email: "123@123",
      title: "목걸이",
      title2: "Tag",
      status: "분실",
      role: "대충 암거나 설명문",
    };

    const users = ref<User[]>([...Array(1).keys()].map(() => testUser));

    return {
      users,
    };
  },
});
</script>
