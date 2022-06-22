<template>
  <div class="container">
    <div class="row m-3">
      <div class="btn-group">
      <button class="iconBtn" title="Dodaj" @click="$router.push('/users/1')">
        <i class="fa fa-user-plus"></i>
      </button>
      <button class="iconBtn" title="Pregledaj"><i class="fa fa-user"></i></button>
      <button class="iconBtn" title="Izmeni"><i class="fa fa-user-md"></i></button>
      <button class="iconBtn" title="Obrisi"><i class="fa fa-user-times"></i></button>
      </div>
    </div>
    <div class="row m-3">
      <vue-table-lite
          :total= 5
          :columns="columns"
          :rows="rows"
          :headerClasses="headerClasses"
      ></vue-table-lite>
    </div>
  </div>
</template>

<script>
import VueTableLite from "vue3-table-lite";
import axios from "axios";

export default {
  name: 'UsersList',
  components: { VueTableLite },
  data() {
    return {
      headerClasses: ["bg-gray"],
      columns: [
        {
          label: 'ID',
          field: 'id',
          width: '3%',
          sortable: true,
          isKey: true,
        },
        {
          label: 'Ime',
          field: 'first_name',
          width: '10%',
          sortable: true,
        },
        {
          label: 'Prezime',
          field: 'last_name',
          width: '10%',
          sortable: true,
        },
        {
          label: 'Email',
          field: 'email',
          width: '10%',
        }
      ],
    rows: [],
    }
  },
  async created() {
    await axios.get('/users/list').then((response) => {
      if (response.data.Data === "") {
        this.errorMsg = "Error getting list!";
        return;
      }
      console.log(response.data.Data);
      this.rows = JSON.parse(response.data.Data);
    }, (error) => {
      console.log(error);
    });
  }
  }
</script>
