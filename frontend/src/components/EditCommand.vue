<template>
  <div id="EditCommand" @click="hideFrame" ref="container">
    <form v-on:submit.prevent="submit">
      <h1>Editer une commande</h1>

      <div>
        <label for="name">Commande: </label>
        <input type="text" id="name" v-model="name" required :disabled="state === 'UPLOADING'"/>

        <label for="help">Help text: </label>
        <input type="text" id="help" v-model="help" required :disabled="state === 'UPLOADING'"/>

        <label for="resp">Description:</label>
        <textarea id="resp" rows="5" v-model="resp" required :disabled="state === 'UPLOADING'"/>

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
  name: 'EditCommand',
  props: ['hide', 'Command'],
  data: function(){
      return {
        state: 'none',
        name: this.Command.name,
        help: this.Command.help,
        resp: this.Command.resp,
      }
  },
  methods: {
    submit() {
      const currThis = this;

      let fd = new FormData();
      fd.append('name', this.name);
      fd.append('help', this.help);
      fd.append('resp', this.resp);

      axios.put('/api/command/' + this.Command.id, fd, { headers: { "Content-Type": "multipart/form-data", "Authorization": this.$store.state.User.token } })
        .then((e) => {
          currThis.state = "SUCCESS";
          currThis.$store.commit('editCommand', e.data);
          currThis.hideDialog();
        })
        .catch((e) => {
          currThis.state = "ERROR";
          // @TODO: Display error
          console.log(e)
        })
    },
    hideFrame(e) {
      if (e.target === this.$refs.container)
        this.hideDialog()
    },
    hideDialog(){
      this.name = '';
      this.help = '';
      this.resp = '';
      this.hide();
    }
  }
}
</script>

<style lang="scss" scoped>

#EditCommand {
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