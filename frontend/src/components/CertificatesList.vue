<template>
    <div class="container">
      <div class="row m-2">
        <div class="btn-group">
          <button class="iconBtn" title="Dodaj" @click="$router.push({name: 'CertificateEdit', params: {certificateId: '', action: 'add' }})">
            <i class="fa fa-plus"></i>
          </button>
          <button class="iconBtn" title="Pregledaj" :disabled="selectedCertificate == null" @click="$router.push({name: 'CertificateEdit', params: {certificateId: selectedCertificate.id, action: 'view' }})" >
            <i class="fa fa-eye"></i>
          </button>
          <button class="iconBtn" title="Izmeni" :disabled="selectedCertificate == null" @click="$router.push({name: 'CertificateEdit', params: {certificateId: selectedCertificate.id, action: 'update' }})" >
            <i class="fa fa-pencil">
            </i>
          </button>
          <button class="iconBtn" title="Obrisi" :disabled="selectedCertificate == null" @click="deleteCertificate">
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
          <div class="collapse multi-collapse border" style="font-size: 0.7em" id="filter">
            <div class="row">
            </div>
      </div>
      <div class="row mt-2">
        <vue-table-lite
            @row-clicked="selectCertificate"
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
      name: 'CertificatesList',
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
              label: 'Marka',
              field: 'brand',
              width: '8%',
            },
            {
              label: 'Tip',
              field: 'type_vehicle',
              width: '6%',
            },
            {
              label: 'Varijanta',
              field: 'variant',
              width: '10%',
            },
            {
              label: 'Verzija',
              field: 'version_vehicle',
              width: '16%',
            },
            {
              label: 'Komercijalna oznaka',
              field: 'commercial_name',
              width: '12%',
            },
            {
              label: 'Kategorija vozila',
              field: 'category',
              width: '12%',
            },
            {
              label: 'Procenjena god. proiz.',
              field: 'estimated_production_year',
              width: '12%',
            },
            {
              label: 'Oznaka motora',
              field: 'engine_code',
              width: '10%',
            },
            {
              label: 'Zapremina motora',
              field: 'engine_capacity',
              width: '12%',
            }
          ],
          rows: [],
          selectedCertificate: null,
          users: [],
          isLoading: false,
          filterObject: {type_vehicle: '', variant: '', version_vehicle: '', estimated_production_year: 0, engine_code: '', engine_capacity: 0, engine_power: 0, fuel: '' },
          totalCount: 0
        }
      },
      methods: {
        selectCertificate(rowData) {
          this.selectedCertificate = rowData;
        },
        getDate(date) {
          return date.split('T')[0];
        },
        async countCertificates() {
          await axios.post('/certificate/count', JSON.stringify(this.filterObject)).then((response) => {
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
          await axios.post('/certificate/list?skip=' + offset + '&take=' + limit, JSON.stringify(this.filterObject)).then((response) => {
            if (response.data === null || response.data.Status === 'error') {
              notie.alert({
                type: 'error',
                text: 'Greska: ' + response.data.ErrorMessage,
                position: 'bottom',
              })
              return;
            }
            this.rows = JSON.parse(response.data.Data);
          }, (error) => {
            notie.alert({
              type: 'error',
              text: "Greska: " + error,
              position: 'bottom',
            })
          });

          await this.countCertificates();

          this.isLoading = false;
        },
        async deleteCertificate() {
          if (confirm("Da li ste sigurni da zelite da obrisete sertifikat?")) {
            await axios.get('/certificate/delete/' + this.selectedCertificate.id).then((response) => {
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
    created() {
        this.doSearch(0, 10);
      },
    mounted() {
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
