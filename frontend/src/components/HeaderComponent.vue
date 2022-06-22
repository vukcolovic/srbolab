<template>
  <nav class="navbar navbar-dark bg-primary navbar-expand-lg">
    <div class="container-fluid">
      <a class="navbar-brand" href="#"><img src="./../assets/img.png"></a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
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
            <router-link class="nav-link" to="/login">Korisnicka podrska</router-link>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>