<template>
  <nav class="bg-gray-800">
    <div class="max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
      <div class="relative flex items-center justify-between h-16">
        <div class="absolute inset-y-0 left-0 flex items-center sm:hidden">
          <!-- Mobile menu button-->
          <button @click="showMenu = !showMenu" type="button"
                  class="inline-flex items-center justify-center p-2 rounded-md text-gray-400 hover:text-white hover:bg-gray-700 focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
                  aria-controls="mobile-menu" aria-expanded="false">
            <span class="sr-only">Open main menu</span>
            <!--
              Icon when menu is closed.

              Heroicon name: outline/menu

              Menu open: "hidden", Menu closed: "block"
            -->
            <svg :class="showMenu ? 'hidden' : 'block'" class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none"
                 viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16M4 18h16"/>
            </svg>
            <!--
              Icon when menu is open.

              Heroicon name: outline/x

              Menu open: "block", Menu closed: "hidden"
            -->
            <svg :class="showMenu ? 'block' : 'hidden'" class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none"
                 viewBox="0 0 24 24" stroke="currentColor" aria-hidden="true">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"/>
            </svg>
          </button>
        </div>
        <div class="flex-1 flex items-center justify-center sm:items-stretch sm:justify-start">
          <div class="flex-shrink-0 flex items-center">
            <img class="block lg:hidden h-12 w-auto" src="~/assets/icons/hat.svg" alt="DSuite">
            <img class="hidden lg:block h-12 w-auto" src="~/assets/icons/hat.svg" alt="Dsuite">
          </div>
          <div class="hidden sm:block sm:ml-6">
            <div class="flex space-x-4">
              <!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->

              <nuxt-link to="/"
                         class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">
                Search
              </nuxt-link>


              <span class="relative inline-flex rounded-md shadow-sm">
                <nuxt-link to="/live"
                           class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">Live</nuxt-link>

                <span class="flex absolute h-3 w-3 top-0 right-0 -mt-1 -mr-1">
                <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-red-400 opacity-75"></span>
                <span class="relative inline-flex rounded-full h-3 w-3 bg-red-500"></span>

                  </span>
              </span>

              <nuxt-link to="/history"
                         class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">
                History
              </nuxt-link>


              <nuxt-link to="/help"
                         class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">
                Help
              </nuxt-link>


            </div>
          </div>
        </div>
        <div class="absolute inset-y-0 right-0 flex items-center pr-2 sm:static sm:inset-auto sm:ml-6 sm:pr-0">


<!--          <button-->
<!--            class="bg-gray-800 p-1 rounded-full text-gray-400 hover:text-white focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white">-->
<!--            <span class="sr-only">View notifications</span>-->
<!--            &lt;!&ndash; Heroicon name: outline/bell &ndash;&gt;-->
<!--            <svg class="h-6 w-6" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"-->
<!--                 stroke="currentColor" aria-hidden="true">-->
<!--              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"-->
<!--                    d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"/>-->
<!--            </svg>-->
<!--          </button>-->

          <!-- Profile dropdown -->
          <div class="ml-3 relative">
            <div>
              <button type="button" @click="showProfile = !showProfile"
                      class="bg-gray-800 flex text-sm rounded-full focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-offset-gray-800 focus:ring-white"
                      id="user-menu-button" aria-expanded="false" aria-haspopup="true">
                <span class="sr-only">Open user menu</span>
                <img class="h-8 w-8 rounded-full" src="~/assets/icons/guy.svg" alt="">
              </button>
            </div>

            <transition
              enter-active-class="transition ease-out duration-100"
              enter-class="transform opacity-0 scale-95"
              enter-to-class="transform opacity-100 scale-100"
              leave-active-class="transition ease-in duration-75"
              leave-class="transform opacity-100 scale-100"
              leave-to-class="transform opacity-0 scale-95"
            >

              <div v-if="showProfile"
                   class="origin-top-right absolute right-0 mt-2 w-48 rounded-md shadow-lg py-1 bg-white ring-1 ring-black ring-opacity-5 focus:outline-none"
                   role="menu" aria-orientation="vertical" aria-labelledby="user-menu-button" tabindex="-1">
                <!-- Active: "bg-gray-100", Not Active: "" -->
                <nuxt-link to="profile" class="block px-4 py-2 text-sm text-gray-700 bg-gray-50" role="menuitem" tabindex="-1"
                   id="user-menu-item-0">Your Profile</nuxt-link>
<!--                <a href="#" class="block px-4 py-2 text-sm text-gray-700" role="menuitem" tabindex="-1"-->
<!--                   id="user-menu-item-1">Settings</a>-->
                <a href="" @click="logout" class="block px-4 py-2 text-sm text-gray-700" role="menuitem" tabindex="-1"
                   id="user-menu-item-2">Sign out</a>
              </div>

            </transition>

          </div>
        </div>
      </div>
    </div>

    <!-- Mobile menu, show/hide based on menu state. -->
    <div v-if="showMenu" class="sm:hidden" id="mobile-menu">
      <div class="px-2 pt-2 pb-3 space-y-1">
        <!-- Current: "bg-gray-900 text-white", Default: "text-gray-300 hover:bg-gray-700 hover:text-white" -->

        <nuxt-link to="/"
                   class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">
          Search
        </nuxt-link>


        <span class="relative inline-flex rounded-md shadow-sm">
                <nuxt-link to="/live"
                           class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">Live</nuxt-link>

                <span class="flex absolute h-3 w-3 top-0 right-0 -mt-1 -mr-1">
                <span class="animate-ping absolute inline-flex h-full w-full rounded-full bg-red-400 opacity-75"></span>
                <span class="relative inline-flex rounded-full h-3 w-3 bg-red-500"></span>

                  </span>
              </span>

        <nuxt-link to="/history"
                   class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">
          History
        </nuxt-link>


        <nuxt-link to="/help"
                   class="text-gray-300 hover:bg-gray-700 hover:text-white block px-3 py-2 rounded-md text-base font-medium">
          Help
        </nuxt-link>


      </div>
    </div>
  </nav>
</template>

<script>
import {toast} from "../helpers/toaster";

export default {
  name: "navbar",


  data() {
    return {
      showProfile: false,
      showMenu: false
    }
  },
  methods: {
    async logout() {
      try {
        await this.$auth.logout()

        let that = this
        await toast(that, "Logged Out !", "success", "top-right", "1500")

        // we want a full refresh await this.$router.push("/")
        await setTimeout(() => location.href = "/", 2000)

      }

        // logout failed ?
      catch (e) {
        console.log(e)
      }
    },
  }
}
</script>

<style scoped>

/*a.nuxt-link-active {*/

/*}*/

#user-menu-item-0 {
  @apply bg-gray-50 font-normal
}
a.nuxt-link-exact-active {
  @apply bg-gray-900 font-bold

}

</style>
