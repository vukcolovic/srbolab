<template>
  <nav class="navbar navbar-dark bg-primary navbar-expand-lg">
    <div class="container-fluid">
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse d-flex">
        <a class="navbar-brand" href="#"><img src="./../assets/img.png"></a>
        <ul class="navbar-nav">
          <li class="nav-item">
            <router-link class="nav-link" to="/login">Vozila</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/login">Administracija</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/login">Izvestaji</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/users">Admin Panel</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/support">Korisnicka podrska</router-link>
          </li>
        </ul>
        <ul class="ms-auto">
          <label style="color: white; font-size: 0.7em">{{firstName}} {{lastName}}</label>
          <button class="iconBtn" title="Odjava" @click="logout"><i class="fa fa-sign-out"></i></button>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>

export default {
  name: 'HeaderComponenet',
  data() {
    return {}
  },
  computed : {
    firstName: function () {
      return this.$store.getters.firstName;
    },
    lastName: function () {
      return this.$store.getters.lastName;
    }
  },
  methods: {
    logout() {
      document.cookie = 'jwt=; Path=/; Expires=Thu, 01 Jan 1970 00:00:01 GMT;';
      this.$store.dispatch('setTokenAction', "");
    }
  },
}
</script>