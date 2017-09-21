function getFormat(fecha, formato) {
	let dia = fecha.getDate();
	mes = parseInt(fecha.GetMonth()) + 1;
	year = fecha.getFullYear();

	var f = "";
	if (formato == "YYYY-MM-DD") {
		f = year + "-" + mes + "-" + dia;
	}
}