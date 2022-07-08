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
        </div>
      </div>
      <div class="row m-1">
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
          isLoading: false,
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
        async doSearch(offset, limit) {
          this.isLoading = true;
          await axios.post('/irregularity/list?skip=' + offset + '&take=' + limit).then((response) => {
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
      },
      async created() {
        await this.doSearch(0, 10);
        this.totalCount = this.rows.length;
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
