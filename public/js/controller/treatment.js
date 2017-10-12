(function() {
  angular.module("treatment.controller", []).controller("TreatmentController", [
    "$scope",
    "$http",
    function($scope, $http) {
      $scope.accion = "";
      $scope.accionTD = "";
      $scope.treatments = {};
      $scope.treatmentsDetail = {};
      $scope.patients = {};
      $scope.work = "";
      $scope.works = {};
      $scope.patientSelected = null;

      $scope.treatment = {
        ID: 0,
        patientId: "",
        select: false
      };
      $scope.treatmentDetail = {
        ID: 0,
        workId: 0,
        quantity: 0,
        treatmentId: null,
        precio: 0
      };

      $scope.btnState = true;

      findAllTreatments($scope, $http);
      // getPatientsWhitoutTreatment($scope, $http);
      getAllWorks($scope, $http);
      // findAllTreatmentsDetail($scope, $http);

      $scope.newTreatment = function() {
        $scope.accion = "NUEVO";
        $scope.treatment.ID = 0;
        getPatientsWhitoutTreatment($scope, $http);
        $scope.treatment.patientId = "";
      };

      $scope.cont = 0;

      $scope.beforeDelete = function(d) {
        //   $scope.work = null;
        $scope.work = d;
        //   console.log(w.works.name);
        $scope.cont++;
        console.log($scope.cont);
      };

      $scope.modifyTreatment = function(t) {
        if ($scope.treatment.treatmentId == null) {
          getAllPatients($scope, $http);
        }
        $scope.accion = "MODIFICAR";
        $scope.treatment.ID = t.ID;
        // console.log(t);
        $scope.treatment.patientId = t.patient;
        console.log($scope.treatment);
      };

      $scope.resetWork = function() {
        $scope.work = null;
      };

      $scope.closeModalForm = function() {
        $scope.treatment = null;
        getAllPatients($scope, $http);
        console.log($scope.treatment);
      };

      $scope.execute = function() {
        switch ($scope.accion) {
          case "NUEVO":
            createTreatment($scope, $http);
            break;
          case "MODIFICAR":
            $scope.treatment.patientId = $scope.treatment.patientId.ID;
            updateTreatment($scope, $http);
            break;
        }
      };

      $scope.addTD = function() {
        $scope.accionTD = "AGREGAR";
        clearFields($scope);
        getAllWorks($scope, $http);
      };

      $scope.modifyTD = function() {
        $scope.accionTD = "MODIFICAR";
      };

      $scope.selectWork = function() {
        console.log($scope.treatmentDetail);
      };

      $scope.selectTreatmentRow = function(id) {
        // console.log("ID: " + id);
        $scope.btnState = false;
        for (let i = 0; i < $scope.treatments.length; i++) {
          if ($scope.treatments[i].ID == id.ID) {
            $scope.treatments[i].state = {
              background: "#51a0c7",
              color: "white"
            };
          } else {
            $scope.treatments[i].state = {
              background: "#eceeef",
              color: "black"
            };
          }
        }
        $scope.treatmentsDetail = {};
        //   $scope.treatmentDetail.treatmentId = id;
        //   console.log(id);
        $scope.patientSelected = id.patient;
        //   console.log($scope.patientSelected);
        findTreatmentsDetailByTreatmentId($scope, $http, id.ID);
      };

      $scope.trStyle = {
        color: "#F5F5F5",
        backgroundColor: "#51a0c7"
      };

      $scope.executeTD = function() {
        console.log($scope.treatmentDetail);
        switch ($scope.accionTD) {
          case "AGREGAR":
            createTD($scope, $http);
            break;
          case "MODIFICAR":
            updateTD($scope, $http);
            break;
        }
      };
    }
  ]);

  function validarCamposTD() {}

  function clearFields($scope) {
    $scope.treatmentDetail.workId = 0;
    $scope.treatmentDetail.quantity = 0;
    // $scope.treatmentDetail.work
  }

  function createTD($scope, $http) {
    $http({
      method: "POST",
      url: "api/crud/treatment-detail/",
      headers: "Content-Type: application/json",
      data: {
        ID: 0,
        workId: $scope.treatmentDetail.workId.ID,
        quantity: $scope.treatmentDetail.quantity,
        treatmentId: $scope.treatmentDetail.treatmentId
      }
    })
      .then(
        function success(response) {
          alert(response.data.message);
          findTreatmentsDetailByTreatmentId(
            $scope,
            $http,
            $scope.treatmentDetail.treatmentId.ID
          );
          // clearFielsTreatments($scope);
        },
        function error(response) {
          console.log($scope.treatment);
          alert(response.data.message);
        }
      )
      .catch(function fail(response) {
        console.log("Excep: " + response.data.message);
      });
  }

  function createTreatment($scope, $http) {
    $http({
      method: "POST",
      url: "api/crud/treatment/",
      headers: "Content-Type: application/json",
      data: {
        ID: 0,
        patientId: $scope.treatment.patientId.ID
      }
    })
      .then(
        function success(response) {
          alert(response.data.message);
          $scope.patients = {};
          findAllTreatments($scope, $http);
          findLastTreatment($scope, $http);
          console.log($scope.treatment);
          $scope.treatmentDetail.treatmentId = $scope.treatment.ID;
          // clearFielsTreatments($scope);
        },
        function error(response) {
          console.log($scope.treatment);
          alert(response.data.message);
        }
      )
      .catch(function fail(response) {
        console.log("Excep: " + response);
      });
  }

  function findLastTreatment($scope, $http) {
    $http({
      method: "GET",
      url: "api/last-treatment/",
      headers: "Content-Type: application/json"
    })
      .then(
        function success(response) {
          // console.log(response.data);
          $scope.treatment = response.data;
        },
        function error(response) {
          alert(response.message);
        }
      )
      .catch(function fail(response) {
        console.log("Excep: " + response);
      });
  }

  function updateTreatment($scope, $http) {
    $http({
      method: "PUT",
      url: "api/crud/treatment/",
      headers: "Content-Type: application/json",
      data: {
        ID: $scope.treatment.ID,
        patientId: $scope.treatment.patientId
      }
    })
      .then(
        function success(response) {
          alert(response.data.message);
          $scope.patients = {};
          findAllTreatments($scope, $http);
        },
        function erroe(response) {
          alert(response.data.message);
        }
      )
      .catch(function fail(response) {
        console.log("Excep: " + response);
      });
  }

  function getAllPatients($scope, $http) {
    $http({
      method: "GET",
      url: "api/patients/",
      headers: "Content-Type: application/json"
    })
      .then(
        function success(response) {
          $scope.patients = response.data;
          // console.log(response.data);
        },
        function error(response) {
          alert(response.message);
        }
      )
      .catch(function fail(response) {
        console.log("Excep: " + response);
      });
  }

  function getPatientsWhitoutTreatment($scope, $http) {
    $http({
      method: "GET",
      url: "api/patients-whitout-treatment/",
      headers: "Content-Type: application/json"
    })
      .then(
        function success(response) {
          $scope.patients = response.data;
          // console.log(response.data);
        },
        function error(response) {
          alert(response.message);
        }
      )
      .catch(function fail(response) {
        console.log("Excep: " + response);
      });
  }

  function getAllWorks($scope, $http) {
    $http({
      method: "GET",
      url: "api/works/",
      headers: "Content-Type: application/json"
    })
      .then(
        function success(response) {
          $scope.works = response.data;
          // console.log(response.data);
        },
        function error(response) {
          alert(response.message);
        }
      )
      .catch(function fail(response) {
        console.log("Excep: " + response);
      });
  }

  function findAllTreatments($scope, $http) {
    $http({
      method: "GET",
      url: "api/treatments/",
      headers: "Content-Type: application/json"
    })
      .then(
        function success(response) {
          $scope.treatments = response.data;
          console.log(response.data);
        },
        function error(response) {
          alert(response.message);
        }
      )
      .catch(function fail(response) {
        console.log("Excep: " + response);
      });
  }

  function findAllTreatmentsDetail($scope, $http) {
    $http({
      method: "GET",
      url: "api/treatments-detail/",
      headers: "Content-Type: application/json"
    })
      .then(
        function success(response) {
          console.log(response.data);
          $scope.treatmentsDetail = response.data;
        },
        function error(response) {
          alert(response.message);
        }
      )
      .catch(function fail(response) {
        console.log("Excep: " + response);
      });
  }

  function findTreatmentsDetailByTreatmentId($scope, $http, id) {
    $http({
      method: "GET",
      url: "api/treatments-detail/" + id,
      headers: "Content-Type: application/json"
    })
      .then(
        function success(response) {
          $scope.treatmentsDetail = response.data;
          console.log(response.data);
          console.log($scope.treatmentDetail);
        },
        function error(response) {
          alert(response.message);
        }
      )
      .catch(function fail(response) {
        console.log("Excep: " + response);
      });
  }

  function clearFielsTreatments($scope) {
    $scope.treatment.ID = 0;
    $scope.treatment.patientId = 0;
  }
})();
