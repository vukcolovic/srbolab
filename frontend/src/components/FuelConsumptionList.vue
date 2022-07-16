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
      </div>
    </div>
<!--    <div class="collapse multi-collapse border" style="font-size: 0.7em" id="filter">-->
<!--      <div class="row">-->
<!--        <div class="col-2 m-2">-->
<!--          <input v-model="filterObject.subject" placeholder="Izvestaj" />-->
<!--        </div>-->
<!--        <div class="col-2 m-2">-->
<!--          <div class="form-check">-->
<!--            <input type="radio" class="form-check-input" id="radio1" name="checkedSubject" value="" v-model="filterObject.checked" checked>Sve-->
<!--            <label class="form-check-label" for="radio1"></label>-->
<!--          </div>-->
<!--          <div class="form-check">-->
<!--            <input type="radio" class="form-check-input" id="radio2" name="checkedSubject" value="true" v-model="filterObject.checked">Ispravljeno-->
<!--            <label class="form-check-label" for="radio2"></label>-->
<!--          </div>-->
<!--          <div class="form-check">-->
<!--            <input type="radio" class="form-check-input" id="radio3" name="checkedSubject" value="false" v-model="filterObject.checked">Neispravljeno-->
<!--            <label class="form-check-label"></label>-->
<!--          </div>-->
<!--        </div>-->
<!--        <div class="col-2 mt-2">-->
<!--          <div class="mb-1">-->
<!--            <label for="datumOd" style="margin-right: 5px">Datum od:</label>-->
<!--            <input type="date" id="datumOd" name="datumOd" v-model="filterObject.date_from" />-->
<!--          </div>-->
<!--          <div>-->
<!--            <label for="datumDo" style="margin-right: 5px">Datum do:</label>-->
<!--            <input type="date" id="datumDo" name="datumDo" v-model="filterObject.date_to" />-->
<!--          </div>-->
<!--        </div>-->
<!--        <div class="col-2 mt-2">-->
<!--          <vue-single-select-->
<!--              name="level"-->
<!--              placeholder="Nivo"-->
<!--              @input="(selected) => filterObject.irregularity_level = selected"-->
<!--              v-model="filterObject.irregularity_level"-->
<!--              :options="irregularityLevels"-->
<!--              option-label="code"-->
<!--          ></vue-single-select>-->
<!--        </div>-->
<!--        <div class="col-2 mt-2">-->
<!--          <vue-single-select-->
<!--              name="Ispitivac"-->
<!--              placeholder="Ispitivac"-->
<!--              @input="(selected) => filterObject.controller = selected"-->
<!--              v-model="filterObject.controller"-->
<!--              :options="users"-->
<!--              option-label="first_name"-->
<!--          ></vue-single-select>-->
<!--        </div>-->
<!--      </div>-->
<!--    </div>-->
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
      // filterObject: {date_from: null, date_to: null, subject: '', irregularity_level: null, controller: null, checked: '' },
      totalCount: 0
    }
  },
  methods: {
    selectFuelConsumption(rowData) {
      this.selectedFuelConsumption = rowData;
    },
    getDate(date) {
      return date.split('T')[0];
    },
    // async countselectedFuelConsumptions() {
    //   await axios.post('/fuel-consumption/count', JSON.stringify(this.filterObject)).then((response) => {
    //     if (response.data === null || response.data.Status === 'error') {
    //       notie.alert({
    //         type: 'error',
    //         text: 'Greska: ' + response.data.ErrorMessage,
    //         position: 'bottom',
    //       })
    //       return;
    //     }
    //     this.totalCount = response.data.Data;
    //   }, (error) => {
    //     notie.alert({
    //       type: 'error',
    //       text: "Greska: " + error,
    //       position: 'bottom',
    //     })
    //   });
    // },
    async doSearch(offset, limit) {
      this.isLoading = true;
      await axios.post('/fuel-consumption/list?skip=' + offset + '&take=' + limit).then((response) => {
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

      // await this.countIrregularities();

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
::v-deep(.vtl-paging-page-label) {
  font-size: 12px;
  padding: 5px;
}
::v-deep(.vtl-paging-pagination-page-link) {
  font-size: 12px;
  padding: 5px;
}
</style>
