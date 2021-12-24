<template>
  <div>

    <div class="flex justify-center items-center">
      <img class="h-14 w-14 mr-2" src="~/assets/icons/suspect.svg">
      <h1 class="text-center text-3xl text-draculaGreen py-4">Association</h1>

    </div>
    <div>

      <div v-if="$fetchState.pending">


        <!-- skeleton loading in 1 line :P  -->
        <div v-for="_ in 5" class="flex justify-center py-4">
          <div class=" bg-gray-400 w-5/6 h-10  rounded-lg animate-pulse"></div>
        </div>

      </div>

      <p v-else-if="$fetchState.error">Error while fetching associations detail please retry later</p>
      <div v-else>


        <accordion title="E-mails" v-if="emails.length > 0 && type !=='email' ">

          <div v-for="email  in emails">

            <span class="px-6" >{{ email }}</span>

          </div>

        </accordion>

        <accordion title="Names" v-if="names.length > 0">

          <div v-for="name in names">
            <span class="px-6">{{ name }}</span>
          </div>

        </accordion>

        <accordion title="Usernames" v-if="usernames.length > 0  && type !=='username'">

          <div v-for="username in usernames">
            <span class="px-6"> {{ username }}</span>
          </div>

        </accordion>

        <accordion title="Addresses" v-if="addresses.length > 0">

          <div v-for="address in addresses">
            <span class="px-6">{{ address }}</span>
          </div>

        </accordion>


        <accordion title="Numbers" v-if="numbers.length > 0 && type !=='number' ">

          <div v-for="number in numbers">
            <span class="px-6"> {{ number }}</span>
          </div>

        </accordion>

        <accordion title="IP addresses" v-if="ips.length > 0  && type !=='ip'">


          <div v-for="ip in ips">
            <span class="px-6"> {{ ip }}</span>
          </div>

        </accordion>


        <accordion title="Domains" v-if="domains.length > 0 && type !=='domain'">

          <div v-for="domain in domains">
            <span class="px-6"> {{ domain }}</span>
          </div>

        </accordion>


      </div>
    </div>

  </div>
</template>

<script>
import Accordion from "./Accordion";

export default {
  components: {Accordion},
  props: ["type", "keyword" , "search_id" ,"history"],
  data() {
    return {
      emails: [],
      names: [],
      usernames: [],
      addresses: [],
      numbers: [],
      ips: [],
      domains: [],
      //passwords: [],
    }
  },
  async fetch() {



    if(this.history){


      let details = await  this.$axios.post("/api/v1/history/search", { s_type:"association", search_id:this.search_id.toString() , keyword:this.keyword})

      this.emails = details.data["emails"]
      this.domains = details.data["domains"]
      this.ips = details.data["ips"]
      this.names = details.data["names"]
      this.numbers = details.data["numbers"]
      this.usernames = details.data["usernames"]
      this.addresses = details.data["addresses"]

      return
    }



    let details = await this.$axios.post("/api/v1/leak/search", {type: this.type, keyword: this.keyword , search_id: this.search_id.toString()})

    this.emails = details.data["emails"]
    this.domains = details.data["domains"]
    this.ips = details.data["ips"]
    this.names = details.data["names"]
    this.numbers = details.data["numbers"]
    this.usernames = details.data["usernames"]
    this.addresses = details.data["addresses"]
    // yeah mate password :P
    // this.passwords = details.data["password"]



  },


}
</script>

<style scoped>

</style>
