<template>
  <form class="home" v-on:submit.prevent="submitForm">
    <h1>Mettre à jours les cours</h1>
    <p>Cette page permet de mettre à jour la commande cours. Téléverser le fichier csv téléchargeable sur votre planning sur l'ENF.</p>

    <input type="file" ref="planning" accept="text/csv"/>

    <span class="error" v-if="Status !== 'NONE' && Status !== 'SUCCESS'">{{Status.length > 0 ? Status : 'Bad request'}}</span>
    <span class="success" v-if="Status === 'SUCCESS'">{{Status}}</span>

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
      let currThis = this;
      currThis.Status = 'NONE';

      let file = this.$refs.planning.files[0];
      let fd = new FormData();
      fd.append('file', file);

      axios({
        method: 'post',
        url: '/api/cours',
        data: fd,
        headers: { 'Content-Type': 'multipart/form-data', 'Authorization': this.$store.state.User.token }
      })
      .then(() => {
        currThis.Status = 'SUCCESS';
      })
      .catch(e => {
        if (e.response)
          currThis.Status = e.response.data;
        console.log(e)
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

.success, .error {
  margin: 1em;
  padding: 1em 3em 1em 3em;
}

.success {
  background: #5cb85c;
  border: 2px solid #4a934a;
  color: #276e37;
}

.error {
  background: #b71c1c;
  border: 2px solid #EF5350;
  color: #ffcdd2;
}

</style>
