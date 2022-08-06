<template>
  <nav class="navbar navbar-dark bg-primary navbar-expand-lg">
    <div class="container-fluid">
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse d-flex">
        <a class="navbar-brand" href="#"><img src="./../assets/img.png"></a>
        <ul class="navbar-nav">
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown"
               aria-expanded="false">
              Vozila
            </a>
            <ul class="dropdown-menu" aria-labelledby="navBarDropDown">
              <li>
                <router-link class="dropdown-item" to="/irregularities">Nepravilnosti</router-link>
              </li>
            </ul>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown"
               aria-expanded="false">
              Administracija
            </a>
            <ul class="dropdown-menu" aria-labelledby="navBarDropDown">
              <li>
                <router-link class="dropdown-item" to="/fuel">Potrosnja goriva</router-link>
              </li>
            </ul>
          </li>
          <li class="nav-item">
            <router-link class="nav-link" to="/certificates">Sertifikati</router-link>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown"
               aria-expanded="false">
              Admin panel
            </a>
            <ul class="dropdown-menu" aria-labelledby="navBarDropDown">
              <li>
                <router-link class="dropdown-item" to="/users">Korisnici</router-link>
              </li>
            </ul>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle" href="#" id="navBarDropDown" role="button" data-bs-toggle="dropdown"
               aria-expanded="false">
              Sifarnici
            </a>
            <ul class="dropdown-menu" aria-labelledby="navBarDropDown">
              <li>
                <router-link class="dropdown-item" to="/irregularity-levels">Nivoi nepravilnosti</router-link>
              </li>
            </ul>
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