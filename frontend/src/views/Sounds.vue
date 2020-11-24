<template>
  <div class="sounds">
    <h1>Soundboard</h1>
    <div id="table">
      <Sound v-for="(s, k) in Sounds" :key="s.id" :Sound="s" :even="k%2 === 0"/>
    </div>
    <button class="discord" type="button" @click="ShowAddSound = true">Ajouter un son</button>
    <AddSound v-if="ShowAddSound" :hide="() => this.ShowAddSound = false"/>
  </div>
</template>

<script>
import {mapState} from "vuex";
import Sound from "@/components/Sound";
import axios from 'axios';
import AddSound from "@/components/AddSound";

export default {
  name: 'Soundboard',
  components: {AddSound, Sound},
  data: function () {
    return {
      ShowAddSound: false
    }
  },
  computed: {
    ...mapState({
      Sounds: state => state.Sounds
    })
  },
  mounted() {
    const commit = this.$store.commit;
    axios.get("/api/sounds", { headers: { "Authorization": this.$store.state.User.token }})
      .then(e => {
        commit('setSounds', e.data)
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
