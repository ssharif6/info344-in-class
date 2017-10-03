$(document).ready(function () {
    //your code here
    $("#form").submit(function(e) {
        console.log('submitted');

        e.preventDefault(); // avoid to execute the actual submit of the form.
        data = getZipcodes($("#input").val());
    });


    function getZipcodes(cityName) {
        $('body').removeAttr("select"); 
        var url = "http://localhost:4000/zips/" + cityName
        $.ajax({
            type: "GET",
            url: url,
            success: function(data) {
                console.log(data);
                arr = data;
                var sel = $('<select>').appendTo('body');
                $(arr).each(function() {
                 sel.append($("<option>").attr('code',this.Code).text(this.Code + ", " + this.State));
                });
            }
        });
    }
});



