<template>
  <div class="sm:container sm:mx-auto">
    <div class="h-10"></div>
    <div class="grid grid-cols-4">
      <div class="font-bold text-xl justify-self-start">
        <p>Services</p>
      </div>
      <div></div>
      <div></div>
      <div class="justify-self-end">
        <KButton appearance="primary">Add New Service</KButton>
      </div>
    </div>
    <div class="h-10"></div>
    <KInput
      type="search"
      placeholder="Search"
      v-model="searchText"
      v-on:change="updateServices"
    ></KInput>
    <div class="h-10"></div>
    <div class="grid grid-cols-4 gap-4">
      <KCard
        v-for="svc in Services"
        :key="svc.ID"
        :title="svc.Name"
        help-text=""
      >
        <template v-slot:body>
          <div class="text-gray-600 w-auto h-40">
            <p class="line-clamp-4">{{ svc.Description }}</p>
          </div>
          <div class="flex gap-3">
            <span
              class="
                bg-transparent
                border border-blue-100
                px-4
                rounded-full
                text-blue-800
              "
            >
              <kLabel> {{ svc.Count }} </kLabel></span
            >
            <span class="font-bold">Versions</span>
          </div>
        </template>
      </KCard>
    </div>
    <div class="h-10"></div>
<div class="sm:container sm:mx-auto flex justify-center space-x-1">
  <KButton apperance="secondary" @click="showBackward">
  <template v-slot:icon>
    <KIcon icon="back" />
  </template>
    </KButton>
    <!-- <a href="#" class=" flex items-center px-4 py-2 text-gray-500 bg-gray-300 rounded-full"> -->
    <!--   <KIcon icon="back"/> -->
    <!-- </a> -->
    <div class="flex items-center" ><p></p></div>
    <!-- <a href="#" class="flex items-center px-4 py-2 text-gray-500 bg-gray-300 rounded-full hover:bg-blue-400 hover:text-white"> -->
    <!--   <KIcon icon="forward"/> -->

      <!-- </a> -->
  <KButton apperance="secondary" @click="showForward">
  <template v-slot:icon>
    <KIcon icon="forward" />
  </template>
    </KButton>

</div>
  </div>
</template>
<script>
import KCard from "@kongponents/kcard";
import KInput from "@kongponents/kinput";
import KButton from "@kongponents/kbutton";
import KIcon from "@kongponents/kicon";
export default {
  name: "ServiceCatalog",
  components: { KCard, KButton, KIcon, KInput },
  props: {},
  data: function () {
    return { Services: [],Links:{}, searchText: ""};
  },
  created() {
    this.updateServices("");
  },
  watch: {
    searchText: function (oldVal, newVal) {
      console.log(oldVal, newVal);
      this.updateServices();
    },
  },
    methods: {
        showBackward: async function(){
            let backwardIndex=this.Links.previous;
      let qs = `sort=+id&limit=12&name[like]=*${encodeURIComponent(this.searchText)}*&${backwardIndex}`;
      let resp = await fetch("http://localhost:8080/services?" + qs).then(
        (response) => response.json()
      );
        this.Services = resp.data;
        this.Links=resp.Links;
        },

        showForward: async function(){
            let forwardIndex=this.Links.next;
      let qs = `sort=+id&limit=12&name[like]=*${encodeURIComponent(this.searchText)}*&${forwardIndex}`;
      let resp = await fetch("http://localhost:8080/services?" + qs).then(
        (response) => response.json()
      );
        this.Services = resp.data;
        this.Links=resp.Links;
        },

    updateServices: async function () {
      let qs = `sort=+id&limit=12&name[like]=*${encodeURIComponent(
        this.searchText
      )}*`;
      let resp = await fetch("http://localhost:8080/services?" + qs).then(
        (response) => response.json()
      );
        this.Services = resp.data;
        this.Links=resp.Links;
    },
  },
};
</script>
