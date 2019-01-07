<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<script>
import PulseLoader from "vue-spinner/src/PulseLoader";
export default {
  name: 'App',
  components: {
    PulseLoader
  }
}
</script>

<style>
body {
  width: 100%;
  height: 100%;
  margin: 0;
}
#app {
  font-family: 'Sarabun', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  width: 100%;
  height: 100%;
}
</style>
