<template>
    <div class="container">
      <div class="row m-2">
        <div class="btn-group">
          <button class="iconBtn" title="Dodaj" @click="$router.push({name: 'IrregularityEdit', params: {irregularityId: '', action: 'add' }})">
            <i class="fa fa-plus"></i>
          </button>
          <button class="iconBtn" title="Pregledaj" :disabled="selectedIrregularity == null" >
            <i class="fa fa-eye"></i>
          </button>
          <button class="iconBtn" title="Izmeni" :disabled="selectedIrregularity == null" >
            <i class="fa fa-pencil">
            </i>
          </button>
          <button class="iconBtn" title="Obrisi" :disabled="selectedIrregularity == null" @click="deleteIrregularity">
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
              <div class="col-2 m-2">
                <input v-model="filterObject.subject" placeholder="Izvestaj" />
              </div>
              <div class="col-2 m-2">
                <div class="form-check">
                  <input type="radio" class="form-check-input" id="radio1" name="checkedSubject" value="" v-model="filterObject.checked" checked>Sve
                  <label class="form-check-label" for="radio1"></label>
                </div>
                <div class="form-check">
                  <input type="radio" class="form-check-input" id="radio2" name="checkedSubject" value="true" v-model="filterObject.checked">Ispravljeno
                  <label class="form-check-label" for="radio2"></label>
                </div>
                <div class="form-check">
                  <input type="radio" class="form-check-input" id="radio3" name="checkedSubject" value="false" v-model="filterObject.checked">Neispravljeno
                  <label class="form-check-label"></label>
                </div>
              </div>
              <div class="col-2 mt-2">
                <div class="mb-1">
                  <label for="datumOd" style="margin-right: 5px">Datum od:</label>
                  <input type="date" id="datumOd" name="datumOd" v-model="filterObject.date_from" />
                </div>
                <div>
                  <label for="datumDo" style="margin-right: 5px">Datum do:</label>
                  <input type="date" id="datumDo" name="datumDo" v-model="filterObject.date_to" />
                </div>
              </div>
              <div class="col-2 mt-2">
                <vue-single-select
                    name="level"
                    placeholder="Nivo"
                    @input="(selected) => filterObject.irregularity_level = selected"
                    v-model="filterObject.irregularity_level"
                    :options="irregularityLevels"
                    option-label="code"
                ></vue-single-select>
              </div>
              <div class="col-2 mt-2">
                <vue-single-select
                    name="Ispitivac"
                    placeholder="Ispitivac"
                    @input="(selected) => filterObject.controller = selected"
                    v-model="filterObject.controller"
                    :options="users"
                    option-label="first_name"
                ></vue-single-select>
              </div>
            </div>
      </div>
      <div class="row mt-2">
        <vue-table-lite
            @row-clicked="selectIrregularity"
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
      name: 'IrregularitiesComponent',
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
              label: 'Predmet',
              field: 'subject',
              width: '10%',
            },
            {
              label: 'Datum',
              field: 'created_at_formated',
              width: '10%',
            },
            {
              label: 'Ustanovljeno',
              field: 'description',
              width: '48%',
            },
            {
              label: 'Napomena',
              field: 'notice',
              width: '15%',
            },
            {
              label: 'Kontrolor',
              field: 'controller_full_name',
              width: '12%',
            },
            {
              label: 'Nivo',
              field: 'level',
              width: '10%',
            },
            {
              label: 'Isp',
              field: 'corrected',
              width: '2%',
              display: function (row) {
                if (row.corrected === true) {
                  return (
                      '<input type="checkbox" onclick="return false;" checked/>'
                  );
                } else {
                  return (
                      '<input type="checkbox" onclick="return false;" />'
                  );
                }
              },
            }
          ],
          rows: [],
          selectedIrregularity: null,
          irregularityLevels: [],
          users: [],
          isLoading: false,
          filterObject: {date_from: null, date_to: null, subject: '', irregularity_level: null, controller: null, checked: '' },
          totalCount: 0
        }
      },
      methods: {
        selectIrregularity(rowData) {
          this.selectedIrregularity = rowData;
        },
        getDate(date) {
          return date.split('T')[0];
        },
        selectLevel(rowData) {
          this.selectedLevel = rowData;
        },
        async countIrregularities() {
          await axios.post('/irregularity/count', JSON.stringify(this.filterObject)).then((response) => {
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
          await axios.post('/irregularity/list?skip=' + offset + '&take=' + limit, JSON.stringify(this.filterObject)).then((response) => {
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
              row.controller_full_name = row.controller.first_name + ' ' +  row.controller.last_name,
              row.corrected_by_full_name = row.corrected_by.first_name + ' ' +  row.corrected_by.last_name,
              row.created_at_formated = this.getDate(row.created_at),
              row.corrected_date_formated = !row.corrected_date.startsWith('0001-01-01') ? this.getDate(row.corrected_date) : "",
              row.level = row.irregularity_level.code
            });
          }, (error) => {
            notie.alert({
              type: 'error',
              text: "Greska: " + error,
              position: 'bottom',
            })
          });

          await this.countIrregularities();

          this.isLoading = false;
        },
        async deleteIrregularity() {
          if (confirm("Da li ste sigurni da zelite da obrisete nepravilnost?")) {
            await axios.get('/irregularity/delete/' + this.selectedIrregularity.id).then((response) => {
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
        async getAllIrregularityLevels() {
          await axios.get('/enumeration/irregularity-levels/all').then((response) => {
            if (response.data === null || response.data.Status === 'error') {
              notie.alert({
                type: 'error',
                text: 'Greska: ' + response.data.ErrorMessage,
                position: 'bottom',
              })
              return;
            }
            this.irregularityLevels = JSON.parse(response.data.Data);
          }, (error) => {
            notie.alert({
              type: 'error',
              text: "Greska: " + error,
              position: 'bottom',
            })
          });
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
        this.getAllIrregularityLevels();
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
