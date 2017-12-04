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
            responsive: false,
            columns: [
                { field: "ID", headerText: "ID" },
                { field: "dateInit", headerText: "Fec. Ingr.", pattern: 'dd/MM/yyyy' },
                { field: "nomApe", headerText: "Nombres y Ape.", filter: true },
                { field: "age", headerText: "Edad" },
                { field: "sex", headerText: "Sexo" },
                { field: "dateNac", headerText: "Fec. Nac." },
                { field: "diabettes", headerText: "Diabts." },
                { field: "hipertension", headerText: "Hiper." }
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
});