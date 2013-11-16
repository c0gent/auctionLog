require(["dojo/dom-attr", "dojo/on", "dojo/dom", "dojo/query", "dojo/domReady!"],
	    function(domAttr, on, dom) {
	 
	        var myObject = {
	            id: "myObject",
	            onClick: function(evt){
	                alert("saleId: " + domAttr.get(evt.target, "data-sale-id"));
	            }
	        };
	        var table = dom.byId("listSalesTable");
	        on(table, ".salesTableEditButton:click", myObject.onClick);
	        on(table, ".salesTableDeleteButton:click", myObject.onClick);

	        //query(".salesTableEditButton").on("click", myObject.onClick);
	        //query(".salesTableDeleteButton").on("click", lang.hitch(myObject, "onClick"));
	 
	});