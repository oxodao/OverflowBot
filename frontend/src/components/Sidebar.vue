<template>
  <nav>
    <router-link to="/">Home</router-link>
    <router-link to="/cours">Cours</router-link>
    <router-link to="/sound">Soundboard</router-link>
  </nav>
</template>

<script>
export default {
  name: 'Sidebar',
}
</script>

<style lang="scss" scoped>
nav {
  background: #23272a;

  a {
    display: block;
    width: 100%;
    height: 3em;
    line-height: 3em;

    color: white;
    text-decoration: none;

    &:hover {
      background: darken(#23272a, 5%);
    }

    &.router-link-exact-active {
      background: lighten(#23272a, 5%);
    }
  }
}
</style>
