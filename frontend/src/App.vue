<template>
  <template v-if="!auth">
    <LoginComponent></LoginComponent>
  </template>
  <template v-else>
    <HeaderComponent></HeaderComponent>
    <router-view></router-view>
    <FooterComponent></FooterComponent>
  </template>
</template>

<script>
import HeaderComponent from "./components/HeaderComponent.vue";
import FooterComponent from "./components/FooterComponent.vue";
import LoginComponent from "@/components/LoginComponent";
import {useStore} from 'vuex';
import {computed} from 'vue';

export default {
  name: 'App',
  components: {
    LoginComponent,
    HeaderComponent,
    FooterComponent,
  },
  setup() {
    const store = useStore();

    const auth = computed(() => store.getters.isAuthenticated)
    return {
      auth
    }
  }
}
</script>

<style>
</style>
