<template>
  <div class="home">
    <h1>Soundboard</h1>
    <div id="table">

    </div>
    <button class="discord" type="button">Ajouter un son</button>
  </div>
</template>

<script>
export default {
  name: 'Soundboard',
}
</script>

<style lang="scss" scoped>

#table {
  flex: 1;
  width: 100%;
  table {
    width: 100%;
  }
}

</style>
