curentRightClickMenuObj=""
 curentRightClickNoteTitleObj=""
 curentOldTitle=""
 curentSelectCateid1="";


 		$(".lefttablabel").each(function(){
			    $(this).click(function(){
			      	$(this).parent().next().toggle();
			    })
    });
 
function loadTreeContextMenu(){

$(".leftmenu li").each(function(){
			    $(this).bind('contextmenu',function(e){
                e.preventDefault();
				curentRightClickMenuObj = $(this);
                $('#mm').menu('show', {
                    left: e.pageX,
                    top: e.pageY
                });
            });
    });
 		
}


function loadTreeClick(){

$(".notecate01").each(function(){
			    $(this).click(function(){ 
			    	var cateid = $(this).attr("catid1");
			    		curentSelectCateid1 = cateid ;

			    		$(".notecate01 a").removeClass("notecate01Select");
			    		$(this).find("a").addClass("notecate01Select");
			    			loadNotes(curentSelectCateid1 , 0);

			    				    })
    });
 		
}



function loadNoteListContextMenu(){

$(".menu2row").each(function(){
			    $(this).bind('contextmenu',function(e){
                e.preventDefault();
				 curentRightClickNoteTitleObj = $(this);
                $('#rightmenuNote').menu('show', {
                    left: e.pageX,
                    top: e.pageY
                });
            });
    });
 		
}

function resetNote()
{
	$("#notetitle").val("");
	UE.getEditor('contentedit').setContent('', false);
	$("#noteid").val("");


}

function NewNote()  {
	resetNote();
}

function    loadNote1(obj)
{
	$(".menu2row").removeClass("menu2rowSelect");
obj.addClass("menu2rowSelect");
var id=obj.attr("artid");


var t = obj.find(".menu2rowtitle").html();

 $("#notetitle").val(t);
 curentOldTitle =t ; 
$("#noteid").val(id);
   $.post("/appser/?act=loadone", { id:id },
   function(data){
     var datac =  data.content ;
	 
 
UE.getEditor('contentedit').setContent(datac, false);


   } ,  "json" );
 
}
	

function addLeftcate(condiv , title ,id , id2)
{
		var str="<li class='notecate01' catid1='"+id +"' cateid2='"+ id2+"'><a >"+ title +"<span>3</span></a></li>";
$(condiv).append(str);
}

$("#menutool_add").click(function(){
//$("#addNotebookPlus").parent().next().show();
 
$("#divinputnewcate").show();
 $("#inputnewcate").show();
 $("#inputnewcate").focus();
})



$("#divinputnewcate").hide();

function LoadLeftTree(){
	$.post("/appser/?act=load", { modname: "notetree" },
	   function(datas){
		  $("#notemenu .notecate01").remove();
		for(i=0 ;i<datas.data.length;i++)
		 {
			var l=datas.data[i].title ;
			var id =datas.data[i].id ;
			
		 	addLeftcate("#notemenu",l,id  ,datas.data[i].cateid2);
		 }

		 loadTreeContextMenu();
		loadTreeClick();
	   } ,  "json" );
}

function AddNoteList(data)
{
var 	str="";
	str+="<div class='menu2row' artid='"+ data.id +"'>";
	str+="<div class='menu2rowtitle'>" + data.title  + "</div>";
	str+="</div>";

	
	$("#menu2list").append(str);
}

function loadNotes(id1 , id2)
{
		$.post("/appser/?act=load", { modname: "notes",cateid:id1 },
   function(datas){
	  $(".menu2row").remove();
	for(i=0 ;i<datas.data.length;i++)
	 {
	 
		 AddNoteList(datas.data[i]);	
	
	   
	 }
	 if(datas.data.length ==0)
	 	{ 
	  			$.messager.show({
				title:'提醒',
				msg:'此类别暂无内容.',
				showType:'show',
				timeout:2000,
				});
	 	}
	 loadNoteListContextMenu();
	 	 $(".menu2row").each(function(){
			    $(this).click(function(){
			       loadNote1($(this));
			    })
    });
   } ,  "json" );
   
}



function saveNote()
{
	var t = $('#notetitle').val();
	var c= UE.getEditor('contentedit').getContent();
	var id = $("#noteid").val();
	  $.post("/appser/?act=save", {id:id , title: t , modname: "notes" ,content: c, cateid:curentSelectCateid1, cateid2:0},
   function(data){
if(data.id > 0)
	   {
			$('#noteid').val(data.id);
			
	   }
	   if(t!=curentOldTitle) //标题改了 ，重新加载一列表。
	   {
			loadNotes(curentSelectCateid1 , 0);
			   curentOldTitle=t ;
	   }
   // addLeftcate("#notebookList",catename,data.id   ,0);
   } ,  "json" );

}

UE.getEditor('contentedit');
$('#notetitle').blur(function(){
saveNote();
});

loadNotes(0 , 0);

UE.getEditor('contentedit').addListener('blur',function(editor){saveNote();});
 function DeleteByID(id)
 {
	if (id<1) return ;
	 $.get("/appser/?act=delete&id=" + id  );
 }

 function doDeleteNote(){
	var nid =  curentRightClickNoteTitleObj.attr("artid") ;
	if(nid>0)
	 {
		DeleteByID(nid);
		curentRightClickNoteTitleObj.remove();
	 }
 }

 LoadLeftTree();

 $('#inputnewcate').blur(function() { 
var catename =  $("#inputnewcate").val();
 
 
  if (catename.length > 1)
  {    
	  $.post("/appser/?act=save", { title: catename , modname: "notetree" },
   function(data){
   	 LoadLeftTree();
   } ,  "json" );

  }

	$("#divinputnewcate").hide();

 

})


 function EditCateName()
 { 
 	var cid = curentRightClickMenuObj.attr("catid1");
 	$.messager.prompt('栏目编辑', '请输入新栏目名', function(r){
				if (r.length>1){
					  $.post("/appser/?act=save", {id:cid , title: r, modname:'notetree'  },
					   function(data){
					      LoadLeftTree();
					   // addLeftcate("#notebookList",catename,data.id   ,0);
					   } ,  "json" );
				}
			});
 }