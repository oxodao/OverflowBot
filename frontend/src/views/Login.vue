<template>
  <div id="login">
    <h1>Overflow's Discord Bot</h1>
    <a class="discord" href="/api/auth">Login with Discord</a>
  </div>
</template>

<script>
export default {
  name: 'Login',
  mounted() {
    const params = this.$route.params;
    if (!!params.id && !!params.username && !!params.token) {
      this.$store.dispatch('setUser', params);
      this.$router.push({name: 'Home' })
    }
  }
}
</script>

<style lang="scss" scoped>
#login {
  width: 100vw;
  height: 100vh;

  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;

  a {
    margin-top: 5em;
  }
}
</style>