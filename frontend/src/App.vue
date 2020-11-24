<template>
  <div id="app">
    <Sidebar id="menu" v-if="isLoggedIn"/>
    <router-view id="page"/>
  </div>
</template>

<script>
import {mapGetters} from "vuex";
import Sidebar from "@/components/Sidebar";

export default {
  name: "App",
  components: {Sidebar},
  computed: {
    ...mapGetters({
      isLoggedIn: 'isLoggedIn',
    })
  }
}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #fff;
  background: #2c2f33;

  width: 100vw;
  height: 100vh;

  display: flex;
  flex-direction: row;

  #menu {
    flex: 0 0 300px;
  }

  #page {
    flex: 1;

    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    padding: 2em;
  }
}

* {
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}

a.discord, button, input[type="submit"] {
  color: white;
  text-decoration: none;

  background: #7289da;
  padding: .5em 1em .25em 1em;
  border: none;
  border-radius: .25em;
}

</style>
