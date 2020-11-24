<template>
  <div :class="'sound ' + (even ? 'even': 'odd')">
    <div class="name">{{Sound.name}}</div>
    <div class="desc">{{Sound.desc}}</div>
    <div class="edit" @click="edit"><img src="@/assets/edit.png" alt="edit"/></div>
    <div class="remove" @click="remove"><img src="@/assets/remove.png" alt="remove"/></div>
    <EditSound v-if="ShowEdit" :Sound="Sound" :hide="() => ShowEdit = false"/>
  </div>
</template>

<script>
import axios from 'axios';
import EditSound from "@/components/EditSound";

export default {
  name: 'Sound',
  components: {EditSound},
  props: [ 'Sound', 'even' ],
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
      const conf = confirm("Voulez vous vraiment supprimer le son '" + this.Sound.name + "' ?");

      if (conf) {
        axios.delete('/api/sound/'+this.Sound.id, { headers: { 'Authorization': this.$store.state.User.token } })
            .then(() => {
              commit('deleteSound', this.Sound.id);
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
.sound {
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

  .name {
    flex: 1;
  }
  .desc {
    flex: 1;
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