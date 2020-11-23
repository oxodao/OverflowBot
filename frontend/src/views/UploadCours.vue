<template>
  <form class="home" v-on:submit.prevent="submitForm">
    <h1>Mettre à jours les cours</h1>
    <p>Cette page permet de mettre à jour la commande cours. Téléverser le fichier csv téléchargeable sur votre planning sur l'ENF.</p>

    <input type="file" id="planning" accept="text/csv"/>

    <span class="error" v-if="Status !== 'NONE' && Status !== 'SUCCESS'">{{Status}}</span>
    <span class="error" v-if="Status === 'SUCCESS'">{{Status}}</span>

    <input class="discord" type="submit"/>
  </form>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Cours',
  data: function() {
    return {
      Status: 'NONE'
    }
  },
  methods: {
    submitForm() {
      this.data.Status = 'NONE';

      let file = this.$refs.plannning.files[0];
      let fd = new FormData();
      fd.append('file', file);

      axios.post('/api/cours', fd, { headers: { 'Content-Type': 'multipart/form-data'}})
      .then(() => {
        this.data.Status = 'SUCCESS';
      })
      .catch(e => {
        if (e.response)
          this.data.Status = e.response.data;
      })
    }
  }
}
</script>

<style lang="scss" scoped>

input {
  margin-top: 2em;
}

p {
  margin-top: 2em;
  width: 350px;
  text-align: justify;
}
</style>
