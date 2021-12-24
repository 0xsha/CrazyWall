<template>
  <div class="text-gray-400">
    <div v-if="$fetchState.pending">
      <Preloader/>
    </div>
    <p v-else-if="$fetchState.error">Error while fetching username detail please retry later</p>
    <div v-else>


      <div class="flex items-center justify-center bg-draculaBackground py-4 mt-2  rounded-lg">

        <img class="h-14 w-14 mr-2" src="~/assets/icons/guy.svg" alt="">


        <h1 class="text-center text-3xl text-green-200 py-4">{{  username }}</h1>

      </div>

      <div class="grid grid-cols-1 gap-4 py-4">

        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">



        <accordion title="Web Accounts" v-if="account_urls.length > 0">

          <div v-for="url  in account_urls">

            <span class="px-6" > <a :href="url" target="_blank"> {{ url}}</a> </span>

          </div>

        </accordion>


          </div>

      </div>

      <div class="grid grid-cols-1 gap-4 py-4">
        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">

          <div v-if="history">
            <Association type="username"  :search_id="this.search_id" :history="this.history"/>
          </div>

          <div v-else>
            <Association type="username" :keyword="this.username" :search_id="this.s_id"/>
          </div>

        </div>


      </div>

    </div>
  </div>

</template>

<script>
import Association from "./Association"
import Preloader from "./Preloader";
import Accordion from "./Accordion";
export default {

  components : {Association,Preloader , Accordion},

  props: ['username', 'history', 'search_id'],


  data(){

    return {
      s_type : "username",
      social_accounts : {},
      account_urls  : [],
      s_id: 0
    }
  },

  async fetch() {


    if (this.history) {

      let details = await this.$axios.post("/api/v1/history/search", {
        s_type: this.s_type,
        search_id: this.search_id.toString()
      })
      this.social_accounts = details.data["socialAccounts"]
      this.s_id = details.data["searchID"]

      let spliced = this.social_accounts.split("\\n")

      let found = []

      for (let i=0;i< spliced.length ; i++) {

        if (spliced[i].indexOf("[+]") !== -1)
        {

          let colIndex = spliced[i].indexOf(": ")
          found.push( spliced[i].substr(colIndex+2))
        }
      }

      this.account_urls = found



      return
    }


    if (!this.history) {

      let response =  await this.$axios.post("/api/v1/username/search", {username: this.username})
      this.s_id = response.data["searchID"]
      this.social_accounts = response.data["socialAccounts"]

      let spliced = this.social_accounts.split("\\n")

      let found = []

      for (let i=0;i< spliced.length ; i++) {

       if (spliced[i].indexOf("[+]") !== -1)
       {

         let colIndex = spliced[i].indexOf(": ")
         found.push( spliced[i].substr(colIndex+2))
       }
      }

      this.account_urls = found

      //console.log(spliced)
    }

  }

}
</script>

<style scoped>

</style>
