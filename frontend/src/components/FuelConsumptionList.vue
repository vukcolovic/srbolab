<template>
  <div class="container">
    <div class="row m-2">
      <div class="btn-group">
        <button class="iconBtn" title="Dodaj" @click="$router.push({name: 'FuelConsumptionEdit', params: {fuelConsumptionId: '', action: 'add' }})">
          <i class="fa fa-plus"></i>
        </button>
        <button class="iconBtn" title="Pregledaj" :disabled="selectedFuelConsumption == null" @click="$router.push({name: 'FuelConsumptionEdit', params: {fuelConsumptionId: selectedFuelConsumption.id, action: 'view' }})" >
          <i class="fa fa-eye"></i>
        </button>
        <button class="iconBtn" title="Izmeni" :disabled="selectedFuelConsumption == null" @click="$router.push({name: 'FuelConsumptionEdit', params: {fuelConsumptionId: selectedFuelConsumption.id, action: 'update' }})" >
          <i class="fa fa-pencil">
          </i>
        </button>
        <button class="iconBtn" title="Obrisi" :disabled="selectedFuelConsumption == null" @click="deleteFuelConsumption">
          <i class="fa fa-trash-o">
          </i>
        </button>
        <button class="iconBtn ms-auto" title="Filter" type="button" data-bs-toggle="collapse" data-bs-target="#filter" aria-expanded="false" aria-controls="filter">
          <i class="fa fa-filter" aria-hidden="true">
          </i>
        </button>
        <button class="iconBtn" title="Trazi" type="button" @click="doSearch(0, 10)">
          <i class="fa fa-search">
          </i>
        </button>
        <button class="iconBtn" title="Ukupan iznos za dati filter" type="button" @click="getSumPrice()">
          <i class="fa fa-money">
          </i>
        </button>
      </div>
    </div>
    <div class="collapse multi-collapse border" style="font-size: 0.7em" id="filter">
      <div class="row">
        <div class="col-3 mt-2">
          <div class="mb-1">
            <label for="datumOd" style="margin-right: 5px">Datum od:</label>
            <input type="date" id="datumOd" name="datumOd" v-model="filterObject.date_from" />
          </div>
          <div>
            <label for="datumDo" style="margin-right: 5px">Datum do:</label>
            <input type="date" id="datumDo" name="datumDo" v-model="filterObject.date_to" />
          </div>
        </div>
        <div class="col-3 m-2">
          <input v-model="filterObject.car_registration" placeholder="Registracija" />
        </div>
        <div class="col-3 mt-2">
          <vue-single-select
              name="Sipao/la"
              placeholder="Sipao/la"
              @input="(selected) => filterObject.poured_by = selected"
              v-model="filterObject.poured_by"
              :options="users"
              option-label="first_name"
          ></vue-single-select>
        </div>
      </div>
    </div>
    <div class="row mt-2">
      <vue-table-lite
          @row-clicked="selectFuelConsumption"
          :columns="columns"
          :rows="rows"
          :total= "totalCount"
          @do-search="doSearch"
          :is-loading="isLoading"
      ></vue-table-lite>
    </div>
  </div>
</template>

<script>
import VueTableLite from "vue3-table-lite";
import axios from "axios";
import notie from 'notie';

export default {
  name: 'FuelConsumptionComponent',
  components: { VueTableLite },
  data() {
    return {
      columns: [
        {
          label: 'ID',
          field: 'id',
          width: '3%',
          isKey: true,
        },
        {
          label: 'Datum',
          field: 'date_consumption',
          width: '10%',
        },
        {
          label: 'Vrsta goriva',
          field: 'fuel_type',
          width: '10%',
        },
        {
          label: 'Kolicina (litara)',
          field: 'liter',
          width: '10%',
        },
        {
          label: 'Iznos (RSD)',
          field: 'price',
          width: '10%',
        },
        {
          label: 'Vozilo (reg. broj)',
          field: 'car_registration',
          width: '10%',
        },
        {
          label: 'Sipao/la',
          field: 'poured_by_fullname',
          width: '10%',
        }
      ],
      rows: [],
      selectedFuelConsumption: null,
      users: [],
      isLoading: false,
      filterObject: {date_from: null, date_to: null, car_registration: '', poured_by: null},
      totalCount: 0,
      sumPrice: 0
    }
  },
  methods: {
    selectFuelConsumption(rowData) {
      this.selectedFuelConsumption = rowData;
    },
    getDate(date) {
      return date.split('T')[0];
    },
    async countFuelConsumptions() {
      await axios.post('/fuel-consumption/count', JSON.stringify(this.filterObject)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        this.totalCount = response.data.Data;
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    },
    async doSearch(offset, limit) {
      this.isLoading = true;
      await axios.post('/fuel-consumption/list?skip=' + offset + '&take=' + limit,  JSON.stringify(this.filterObject)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        this.rows = JSON.parse(response.data.Data);

        this.rows.forEach(row => {
          row.poured_by_fullname = row.poured_by.first_name + ' ' +  row.poured_by.last_name
        });
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });

      await this.countFuelConsumptions();

      this.isLoading = false;
    },
    async deleteFuelConsumption() {
      if (confirm("Da li ste sigurni da zelite da obrisete?")) {
        await axios.get('/fuel-consumption/delete/' + this.selectedFuelConsumption.id).then((response) => {
          if (response.data === null || response.data.Status === 'error') {
            notie.alert({
              type: 'error',
              text: 'Greska: ' + response.data.ErrorMessage,
              position: 'bottom',
            })
          }
        }, (error) => {
          notie.alert({
            type: 'error',
            text: "Greska: " + error,
            position: 'bottom',
          })
        });
      }

      location.reload();
    },
    async getAllUsers() {
      await axios.get('/users/list?skip=0&take=100').then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        this.users = JSON.parse(response.data.Data);
        this.users.forEach(user => user.first_name = user.first_name + ' ' +  user.last_name);
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    },
    async getSumPrice() {
      await axios.post('/fuel-consumption/sum-price', JSON.stringify(this.filterObject)).then((response) => {
        if (response.data === null || response.data.Status === 'error') {
          notie.alert({
            type: 'error',
            text: 'Greska: ' + response.data.ErrorMessage,
            position: 'bottom',
          })
          return;
        }
        this.sumPrice = response.data.Data;
        alert("Ukupan placeni iznos za dati filter je: " + this.sumPrice);
      }, (error) => {
        notie.alert({
          type: 'error',
          text: "Greska: " + error,
          position: 'bottom',
        })
      });
    },
  },
  mounted() {
    this.doSearch(0, 10);
    this.getAllUsers();
  }
}
</script>

<style scoped>
::v-deep(.vtl-table .vtl-thead .vtl-thead-th) {
  font-size: 12px;
}
::v-deep(.vtl-table td),
::v-deep(.vtl-table tr) {
  font-size: 12px;
  padding: 5px;
}
::v-deep(.vtl-paging-info) {
  font-size: 12px;
  padding: 5px;
}
::v-deep(.vtl-paging-count-label),
::v-deep(.vtl-paging-page-label),
::v-deep(.vtl-paging-count-dropdown),
::v-deep(.vtl-paging-page-dropdown){
  font-size: 12px;
  padding: 5px;
}
::v-deep(.vtl-paging-pagination-page-link) {
  font-size: 12px;
  padding: 5px;
}
</style>
