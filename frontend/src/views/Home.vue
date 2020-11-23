<template>
  <div class="home">
    <h1>Overflow Bot</h1>
    <p>Attention, veuillez ne pas appeler le bot s'il est déjà présent sur un serveur. Il faut créer une nouvelle instance complète, le code ne gère pas le multi-serveur. Cela cassera son intégration sur TOUS LES SERVEURS SUR LEQUEL IL EST CONNECTÉ</p>
    <a class="discord" href="/api/call">Appeler le bot sur son serveur</a>
  </div>
</template>

<script>
export default {
  name: 'Home',
}
</script>

<style lang="scss" scoped>
a {
  margin-top: 5em;
}

p {
  margin-top: 2em;
  width: 350px;
  text-align: justify;
}
</style>
