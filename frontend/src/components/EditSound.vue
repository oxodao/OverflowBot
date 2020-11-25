<template>
  <div id="EditSound" @click="hideFrame" ref="container">
    <form v-on:submit.prevent="submit">
      <h1>Editer un son</h1>

      <div>
        <label for="name">Nom / Commande: </label>
        <input type="text" ref="name" id="name" v-model="name" required :disabled="state === 'UPLOADING'"/>

        <label for="desc">Description:</label>
        <textarea id="desc" ref="desc" rows="5" v-model="desc" required :disabled="state === 'UPLOADING'"/>

        <span v-if="state === 'UPLOADING'">Uploading...</span>
        <span class="error" v-if="state === 'ERROR'">An error occured!</span>
      </div>

      <input type="submit" value="Editer"/>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'EditSound',
  props: ['hide', 'Sound'],
  data: function(){
      return {
        state: 'none',
        name: this.Sound.name,
        desc: this.Sound.desc,
      }
  },
  methods: {
    submit() {
      const currThis = this;

      let fd = new FormData();
      fd.append('name', this.name);
      fd.append('desc', this.desc);

      axios.put('/api/sound/' + this.Sound.id, fd, { headers: { "Content-Type": "multipart/form-data", "Authorization": this.$store.state.User.token } })
        .then((e) => {
          currThis.state = "SUCCESS";
          currThis.$store.commit('editSound', e.data);
          currThis.hideDialog();
        })
        .catch(() => {
          currThis.state = "ERROR";
          // @TODO: Display error
        })
    },
    hideFrame(e) {
      if (e.target === this.$refs.container)
        this.hideDialog()
    },
    hideDialog(){
      this.name = '';
      this.desc = '';
      this.hide();
    }
  }
}
</script>

<style lang="scss" scoped>

#EditSound {
  position: absolute;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;

  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  background: rgba(0, 0, 0, .8);

  form {
    width: 300px;
    height: 400px;

    background: #23272a;
    color: white;
    border-radius: 1em;
    border: 1px solid black;

    display: flex;
    flex-direction: column;
    align-items: center;

    h1 {
      margin-top: .5em;
    }

    div {
      margin: 1em 0 1em 0;
      flex: 1;
      width: 80%;

      label, input, textarea {
        display: block;
        width: 100%;
      }

      label:not(:nth-child(1)) {
        margin-top: .5em;
      }

    }

    input[type="submit"] {
      margin-bottom: .5em;
    }

  }
}
</style>
