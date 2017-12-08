$(document).ready(function() {
  // alert("Loading...")
  var url_1 = "http://192.168.1.33/odontosoft-php/api/patient_api.php";
  var url_2 = "http://localhost/odontosoft-php/api/patient_api.php";
  var url_3 = "http://localhost:3030/api/patients/";

  // var patients = null;

  $(function() {
    $("#tblPatients").puidatatable({
      caption: "Lista de Pacientes",
      paginator: {
        rows: 10
      },
      selectionMode: "single",
      responsive: true,
      columns: [
        {
          field: "ID",
          headerText: "ID",
          bodyClass: "cell-content",
          headerClass: "header-table-small"
        },
        {
          field: "date_init",
          headerText: "Fec. Ingr.",
          pattern: "dd/MM/yyyy",
          content: function(p) {
            return setFormatDate(p.date_init);
          },
          headerClass: "header-date",
          filter: true,
          sortable: true
        },
        {
          field: "nom_ape",
          headerText: "Nombres y Ape.",
          filter: true,
          sortable: true,
          headerClass: "clumn-hight"
        },
        {
          field: "address",
          headerText: "Direccón",
          filter: true,
          sortable: true,
          headerClass: "clumn-hight"
        },
        {
          field: "tel_cel",
          headerText: "Telef. / Cel.",
          filter: false,
          sortable: false
        },
        {
          field: "age",
          headerText: "Edad",
          bodyClass: "cell-content",
          headerClass: "header-table-small",
          filter: true,
          sortable: true
        },
        {
          field: "sex",
          headerText: "Sexo",
          bodyClass: "cell-content",
          headerClass: "header-table-small",
          filter: true,
          sortable: true
        },
        {
          field: "date_nac",
          headerText: "Fec. Nac.",
          content: function(p) {
            return setFormatDate(p.date_nac);
          },
          headerClass: "header-date",
          filter: true,
          sortable: true
        },
        {
          field: "diabettes",
          headerText: "Diabts.",
          headerClass: "header-enf",
          bodyClass: "cell-content",
          content: function(p) {
            return getStateEnfermedad(p.diabettes);
          },
          filter: true,
          sortable: true
        },
        {
          field: "hipertension",
          headerText: "Hiper.",
          headerClass: "header-enf",
          bodyClass: "cell-content",
          content: function(p) {
            return getStateEnfermedad(p.hipertension);
          }
        }
      ],
      datasource: function(callback) {
        $.ajax({
          type: "GET",
          url: url_3,
          data: {
            opt: "all"
          },
          dataType: "json",
          context: this,
          success: function(response) {
            let patients;
            for (let i = 0; i < response.length; i++) {
              // console.log(response[i]);
            }
            callback.call(this, response);
          }
        });
      }
    });
  });

  $("#frmPatient").puidialog({
    showEffect: "fade",
    hideEffect: "fade",
    modal: true,
    minimizable: false,
    maximizable: false,
    responsive: true,
    //   minWidth: 400,
    width: 600,
    buttons: [
      {
        text: "GUARDAR",
        icon: "fa-floppy-o",
        click: function() {
          $("#frmPatient").puidialog("hide");
        }
      }
    ]
  });
});

$("#pnlControlPatient").puipanel();

$("#btnAdd").puibutton({
  icon: "fa-plus",
  click: function() {
    $("#frmPatient").puidialog("show");
  }
});

$(function(){
    $("#txtId").puiinputtext();
    $("#txtNomApe").puiinputtext();
});

$("p-datepicker")
  .children("input")
  .puiinputtext();

$("#btnModify").puibutton({
  icon: "fa-pencil"
});

function setFormatDate(fec) {
  return moment(fec).format("DD/MM/YYYY");
}

function getStateEnfermedad(enf) {
  if (enf) {
    return $("<i class='fa fa-check aria-hidden='true'></i>");
  } else {
    return $("<i class='fa fa-times aria-hidden='true'></i>");
  }
}

$("#btnModificar").click(function() {
  alert("Modificando paciente: " + p.nom_ape);
});

function modifyPatient(p) {
  alert("Modificando paciente: " + p.nom_ape);
}

function setButtonStyle() {
  $(".btnUpdate").puibutton({
    icon: "fa fa-pencil"
  });
}

// $('#btnAdd').click(function(){
//     document.getElementById("frmPatient").show();
// });
