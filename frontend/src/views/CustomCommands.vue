<template>
  <div class="sounds">
    <h1>Custom commands</h1>
    <div id="table">
      <CustomCommand v-for="(c, k) in Commands" :key="c.id" :Command="c" :even="k%2 === 0"/>
    </div>
    <button class="discord" type="button" @click="ShowAddCommand= true">Ajouter une commande</button>
    <AddCommand v-if="ShowAddCommand" :hide="() => this.ShowAddCommand = false"/>
  </div>
</template>

<script>
import {mapState} from "vuex";
import CustomCommand from "@/components/Command";
import axios from 'axios';
import AddCommand from "@/components/AddCommand";

export default {
  name: 'CustomCommands',
  components: {CustomCommand, AddCommand },
  data: function () {
    return {
      ShowAddCommand: false
    }
  },
  computed: {
    ...mapState({
      Commands: state => state.Commands
    })
  },
  mounted() {
    const commit = this.$store.commit;
    axios.get("/api/commands", { headers: { "Authorization": this.$store.state.User.token }})
      .then(e => {
        commit('setCommands', e.data)
      })
  }
}
</script>

<style lang="scss" scoped>

.sounds {
  height: 100%;
}

#table {
  flex: 1;
  width: 100%;
  padding: 2em;
  height: 0;
  overflow-y: scroll;
}

button {
  margin-top: 1em;
}

</style>
