<template>
  <div>

    <div class="container mt-12">
      <div class="flex justify-center items-center">
        <div>

          <img class="mt-4 mb-4 rounded-lg w-96"
               src="https://static.vecteezy.com/system/resources/previews/002/218/098/non_2x/desk-of-a-detective-investigator-and-night-shift-worker-vector.jpg"
               alt="">
          <div>

            <div class="flex flex justify-end">
              <input type="text" v-model="keyword" @keyup.enter="search" ref="searchInput"
                     class="relative px-4 py-3  w-96 leading-5 border rounded-md focus:outline-none focus:ring focus:border-blue-400"
                     placeholder="domain, email, phone, name, username, ip"  >
              <img class="absolute  h-10 w-10 pt-2" src="~/assets/icons/loupe.svg" alt="">
            </div>

            <div class="flex justify-center items-center py-4" @click="searchSwitch">
            <span class="dark:text-draculaGreen px-2 font-bold " ref="switchText">Databases</span>
              <div class="w-16 h-6 flex items-center bg-gray-300 rounded-full p-1 duration-300 ease-in-out" :class="{ 'bg-red-600': toggleActive}">
                <div class="bg-white w-4 h-4 rounded-full shadow-md transform duration-300 ease-in-out" :class="{ 'translate-x-10': toggleActive,}"></div>
              </div>
            </div>


            </div>

          </div>
        </div>
      </div>


    </div>



</template>

<script>

import Flag from "../components/Flag";
import {toast} from "../helpers/toaster";
import Association from "../components/Association";
import Phone from "../components/Phone";

export default {
  components: {Phone, Association, Flag},
  data() {
    return {
      keyword: '',
      phone: false,
      toggleActive: false,
    }
  },
  methods: {
    async logout() {
      try {
        await this.$auth.logout()
        location.reload()
      }

        // logout failed ?
      catch (e) {
        console.log(e)
      }
    },
    getResults(type) {

      // this.$auth.$storage.setLocalStorage("type", type)
      // this.$auth.$storage.setLocalStorage("keyword", this.keyword)


      this.$router.push({name : "result", params : {type: type , keyword: this.keyword}})
    },

    searchSwitch(){
      this.toggleActive = !this.toggleActive

      if (this.toggleActive) {
        this.$refs.searchInput.placeholder = "keyword (across social networks and news)"
        this.$refs.switchText.textContent = "Live"
      }
      else {
        this.$refs.searchInput.placeholder = "domain, email, phone, name, username, IP"
        this.$refs.switchText.textContent = "Databases"
      }

    },
    search() {
      //  let res = await this.$axios.post("/api/v1/phone/search", {phone: this.keyword})
      // console.log(res)

      let searchType = this.$refs.switchText.textContent


      if (searchType === "Databases")
      {


        let fullNamePattern = /^(([A-Za-z]+[\-\']?)*([A-Za-z]+)?\s)+([A-Za-z]+[\-\']?)*([A-Za-z]+)?$/
        let phonePattern = /^(?:[0-9] ?){6,14}[0-9]$/
        let userNamePattern = /^[a-zA-Z0-9]+$/;
        let emailPattern = /^[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?$/
        let ipPattern = /^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$/
        let domainPattern = /^(?!:\/\/)([a-zA-Z0-9]+\.)?[a-zA-Z0-9][a-zA-Z0-9-]+\.[a-zA-Z]{2,6}?$/i // insensitive

        //let intValue = parseInt(this.keyword)
        let matchPhone = phonePattern.exec(this.keyword)
        let matchEmail = emailPattern.exec(this.keyword)
        let matchUsername = userNamePattern.exec(this.keyword)
        let matchIP = ipPattern.exec(this.keyword)
        let matchDomain = domainPattern.exec(this.keyword)
        let matchFullName = fullNamePattern.exec(this.keyword)




        if (matchPhone) {


          this.getResults("phone")

        } else if (matchUsername) {


          this.getResults("username")

        } else if (matchEmail) {



          this.getResults("email")

        } else if (matchIP) {



          this.getResults("ip")

        } else if (matchDomain) {



          this.getResults("domain")

        } else if (matchFullName) {


          this.getResults("name")
        }
        else {

          let that = this
          toast(  that,"Invalid input","error" , "top-right" , 3000)

        }

        // no else end the function
        return
      }

      // live search code here
     // this.$auth.$storage.setLocalStorage("keyword", this.keyword)
      this.$router.push({name : "live", params : {keyword: this.keyword}})


    }
  }
}
</script>

<style>

</style>
