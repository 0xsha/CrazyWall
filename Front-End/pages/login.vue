<template>
  <div class="min-h-screen flex items-center justify-center bg-gray-50  py-12 px-4 sm:px-6 lg:px-8">
    <div class="max-w-md w-full space-y-8">
      <div>
        <img class="mx-auto h-48 w-auto" src="~/assets/icons/detective-full.svg" alt="Workflow" />
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900 dark:text-draculaGreen">
          Sign in to your account
        </h2>
        <p class="mt-2 text-center text-sm text-gray-600">

          <span v-if="error" class="font-medium text-indigo-600 dark:text-red-600 hover:text-indigo-500">
          {{errorMsg}}
          </span>
        </p>
      </div>

      <form class="mt-8 space-y-6" @submit.prevent="userLogin">
        <input type="hidden" name="remember" value="true" />
        <div class="rounded-md shadow-sm -space-y-px">
          <div>
            <label for="email-address" class="sr-only">Email address</label>
            <input id="email-address" v-model="login.email" name="email" type="email" autocomplete="email" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Email address" />
          </div>
          <div>
            <label for="password" class="sr-only">Password</label>
            <input id="password" v-model="login.password" name="password" type="password" autocomplete="current-password" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Password" />
          </div>
        </div>
          <div class="text-sm flex justify-between">
            <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">
              Forgot your password?
            </a>

            <NuxtLink to="/register"  class="font-medium text-indigo-600 hover:text-indigo-500">   Need an account?</NuxtLink>

          </div>
        <div>
          <button type="submit" class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white dark:bg-draculaPurple dark:hover:bg-purple-700 bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            <span class="absolute left-0 inset-y-0 flex items-center pl-3">
<!--              <LockOpenIcon class="h-5 w-5 text-indigo-500 group-hover:text-indigo-400" aria-hidden="true" />-->
            </span>
            Sign in
          </button>
        </div>
      </form>
    </div>
  </div>

</template>

<script>
// import { LockOpenIcon } from '@heroicons/vue/solid'

import {toast} from "../helpers/toaster";

export default {

  layout: "auth",
  // components: {
  //   LockOpenIcon,
  // },
  data() {
    return {
      error: false,
      errorMsg : '',
      login: {
        email: '',
        password: ''
      }
    }
  },
  methods: {
    async userLogin() {

      try {


        let response = await this.$auth.loginWith('local', { data: this.login })

        if (response.status === 200) {

          // if this.$auth.strategy.token.status()._status == "VALID"
          // this.$auth.setUserToken(response.data.token)

          let that = this
          await toast(  that,"Logged In !","success" , "top-right" , "2000")

          await this.$router.push("/")

        }

      } catch (err) {
        this.error = true
        this.errorMsg = err.response.data


        let that = this
        await toast(  that,this.errorMsg,"error" ,"top-right" , "2000")

      }
    }
  }
}
</script>
