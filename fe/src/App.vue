<template>
  <div id="app">
    <SayThanks/>
  </div>
</template>

<script>
import SayThanks from './components/SayThanks.vue'

export default {
  name: 'app',
  components: {
    SayThanks
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
