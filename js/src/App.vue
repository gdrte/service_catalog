<template>
  <div id="app">
    <NaviBar/>
    <div class="h-1"></div>
    <ServiceCatalog title="Welcome to Kong" />
  </div>
</template>

<script>
import ServiceCatalog from "./components/ServiceCatalog.vue";
import NaviBar from "./components/NaviBar.vue";

export default {
  name: "App",
  components: {
    ServiceCatalog,NaviBar
  },
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
