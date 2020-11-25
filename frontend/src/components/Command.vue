<template>
  <div :class="'command ' + (even ? 'even': 'odd')">
    <div class="name">{{Command.name}}</div>
    <div class="help">{{Command.help}}</div>
    <pre class="resp">{{Command.resp}}</pre>
    <div class="edit" @click="edit"><img src="@/assets/edit.png" alt="edit"/></div>
    <div class="remove" @click="remove"><img src="@/assets/remove.png" alt="remove"/></div>
    <EditCommand v-if="ShowEdit" :Command="Command" :hide="() => ShowEdit = false"/>
  </div>
</template>

<script>
import axios from 'axios';
import EditCommand from "@/components/EditCommand";

export default {
  name: 'Command',
  components: {EditCommand},
  props: [ 'Command', 'even' ],
  data: function () {
    return {
      ShowEdit: false
    }
  },
  methods: {
    edit() {
      this.ShowEdit = true;
    },
    remove() {
      const commit = this.$store.commit;
      const conf = confirm("Voulez vous vraiment supprimer la commande '" + this.Command.name + "' ?");

      if (conf) {
        axios.delete('/api/command/'+this.Command.id, { headers: { 'Authorization': this.$store.state.User.token } })
            .then(() => {
              commit('deleteCommand', this.Command.id);
            })
            .catch(() => {
              alert("Une erreur est survenue lors de la suppression.");
            })
      }
    }
  }
}
</script>

<style lang="scss" scoped>
.command {
  display: flex;
  flex-direction: row;

  padding: .5em;

  img {
    width: 2em;
    margin-left: 1em;
  }

  div {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
  }
  div, pre{
    padding: .5em;
  }

  .name {
    flex: 1;
  }
  .resp {
    flex: 1;
    text-align: left;
  }

  &.even {
    background: #2f3439;
  }

  &.odd {
    background: #373d42;
  }

  .edit, .remove {
    cursor: pointer;
  }
}
</style>