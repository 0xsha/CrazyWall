<template>


  <div class="min-h-screen flex items-center justify-left  bg-gray-50 dark:bg-gray-900 py-14 px-4 sm:px-6 lg:px-14">
    <div class="max-w-md w-full space-y-8">
      <div>
        <img class="mx-auto h-48 w-auto" src="~/assets/icons/detective-full.svg" alt="DSuite" />
        <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900 dark:text-draculaGreen">
          Create a new account
        </h2>
<!--        <p class="mt-2 text-center text-sm text-gray-600">-->

<!--          <span v-if="error" class="font-medium text-indigo-600 dark:text-red-600 hover:text-indigo-500">-->
<!--          {{errorMsg}}-->
<!--          </span>-->
<!--        </p>-->

        <div v-if="error" class="mt-2 text-center font-medium text-indigo-600 dark:text-red-600 hover:text-indigo-500">
          <div> {{errMsg}}</div>
        </div>
      </div>

      <form class="mt-8 space-y-6" @submit.prevent="registerUser">
        <input type="hidden" name="remember" value="true" />
        <div class="rounded-md shadow-sm -space-y-px">
          <div>
            <label for="username" class="sr-only">User Name</label>
            <input id="username" v-model="register.username" name="username" type="text" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Username" />
          </div>

          <div>
            <label for="email-address" class="sr-only">Email address</label>
            <input id="email-address" v-model="register.email" name="email" type="email" autocomplete="email" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900  focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Email address" />
          </div>
          <div>
            <label for="password" class="sr-only">Password</label>
            <input id="password" v-model="register.password" name="password" type="password" autocomplete="current-password" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Password" />
          </div>

          <div>
            <label for="password_confirm" class="sr-only">Password</label>
            <input id="password_confirm" v-model="register.password_confirm" name="password_confirm" type="password" autocomplete="current-password" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Confirmation" />
          </div>

          <div>
            <label for="invite_code" class="sr-only">Invite code</label>
            <input id="invite_code" v-model="register.invite_code" name="invite_code" type="text" required="" class="appearance-none rounded-none relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-b-md focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 focus:z-10 sm:text-sm" placeholder="Invite Code" />
          </div>

        </div>
        <div class="text-sm flex justify-between">
          <a href="#" class="font-medium text-indigo-600 hover:text-indigo-500">
            Forgot your password?
          </a>

          <NuxtLink to="/login"  class="font-medium text-indigo-600 hover:text-indigo-500">   Already have an account?</NuxtLink>

        </div>
        <div>
          <button type="submit" class="group dark:bg-draculaPurple dark:hover:bg-purple-700 relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500">
            <span class="absolute left-0 inset-y-0 flex items-center pl-3">
<!--              <LockOpenIcon class="h-5 w-5 text-indigo-500 group-hover:text-indigo-400" aria-hidden="true" />-->
            </span>
            Sign up
          </button>
        </div>


      </form>
    </div>
    <div class="flex justify-center">

      <img class=" w-4/5 rounded-lg  " src="https://static.vecteezy.com/system/resources/previews/002/217/917/original/seamless-pattern-of-detective-investigation-vector.jpg" alt="">
    </div>
  </div>

</template>

<script>
import {toast} from "../helpers/toaster";

export default {
  // guest => redirect to home if logged in
  auth: 'guest',
  layout: "auth",
  data() {
    return {
      error: false,
      errMsg: '',
      register: {
        username: '',
        email: '',
        password: '',
        password_confirm:'',
        invite_code: ''

      }
    }
  },

  methods: {
    async registerUser() {

      try {

        let response = await this.$axios.post("/api/v1/auth/register"  , this.register)

        if (response.status === 201) {

          let that = this
          await toast(  that,"Account Created !","success" , "top-right" , "2000")

         await this.$router.push("/")

        }

      } catch (err) {

        // there is an error
        this.error = true
        this.errMsg = err.response.data.Error
        if (err.response.status === 429)
        {

          this.errMsg = err.response.data
        }

      }
    }
  }
}
</script>
