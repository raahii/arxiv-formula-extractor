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

<style lang="scss">
html {
  width: 100%;
  height: 100%;
}
body {
  width: 100%;
  height: 100%;
  margin: 0;
}
#app {
  width: 100%;
  height: 100%;
  font-family: 'Sarabun', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;

  .error {
    font-size: 1.2em;
    line-height: 1.5;
    color: #a94442;
    background-color: #f2dede;

    box-sizing: border-box;
    width: 100%;
    padding: 15px 15px 15px 35px;
    margin: 20px 0;

    border-color: #ebccd1;
    border: 1px solid transparent;
    border-radius: 4px;
  }
  .success{
    font-size: 1.2em;
    line-height: 1.5;
    color: #31708f;
    background-color: #d9edf7;

    box-sizing: border-box;
    width: 100%;
    padding: 15px 15px 15px 35px;
    margin: 20px 0;

    border-color: #bce8f1;
    border: 1px solid transparent;
    border-radius: 4px;
  }
}
</style>
