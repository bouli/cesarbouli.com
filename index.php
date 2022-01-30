<?php 
$meta_image = "http://cesarbouli.com/images/photo.jpg";
$meta_text = "Escute Cesar Bouli.";
$meta_title = "Cesar Bouli";
?><!DOCTYPE html>
<html lang="es">
<head> 
	<!-- Global site tag (gtag.js) - Google Analytics -->
	<script async src="https://www.googletagmanager.com/gtag/js?id=UA-122379890-1"></script>
	<script>
	  window.dataLayer = window.dataLayer || [];
	  function gtag(){dataLayer.push(arguments);}
	  gtag('js', new Date());

	  gtag('config', 'UA-166762733-1');
	</script>
	<script src="https://code.jquery.com/jquery-1.11.3.js"></script>
	<meta charset="UTF-8">
	<link rel="apple-touch-icon" sizes="57x57" href="images/apple-icon-57x57.png">
	<link rel="apple-touch-icon" sizes="60x60" href="images/apple-icon-60x60.png">
	<link rel="apple-touch-icon" sizes="72x72" href="images/apple-icon-72x72.png">
	<link rel="apple-touch-icon" sizes="76x76" href="images/apple-icon-76x76.png">

	<link rel="apple-touch-icon" sizes="114x114" href="images/apple-icon-114x114.png">
	<link rel="apple-touch-icon" sizes="120x120" href="images/apple-icon-120x120.png">
	<link rel="apple-touch-icon" sizes="144x144" href="images/apple-icon-144x144.png">
	<link rel="apple-touch-icon" sizes="152x152" href="images/apple-icon-152x152.png">
	<link rel="apple-touch-icon" sizes="180x180" href="images/apple-icon-180x180.png">

	<link rel="icon" type="image/png" sizes="192x192"  href="images/android-icon-192x192.png">
	<link rel="icon" type="image/png" sizes="32x32" href="images/favicon-32x32.png">
	<link rel="icon" type="image/png" sizes="96x96" href="images/favicon-96x96.png">
	<link rel="icon" type="image/png" sizes="16x16" href="images/favicon-16x16.png">

	<link rel="manifest" href="images/manifest.json">

	<meta name="msapplication-TileColor" content="#000000">
	<meta name="msapplication-TileImage" content="images/ms-icon-144x144.png">
	
	<meta name="theme-color" content="#000000">
	<meta name="description" content="<?php echo $meta_text; ?>" /> 

	<meta property="og:locale" content="pt-br" />
	<meta property="og:url" content="http://cesarbouli.com"> 
	<meta property="og:title" content="<?php echo $meta_title; ?>"> 
	<meta property="og:site_name" content="<?php echo $meta_title; ?>"> 
	<meta property="og:description" content="<?php echo $meta_text; ?>"> 
	<meta property="og:image" content="<?php echo $meta_image; ?>">
	<meta property="og:image:type" content="image/png">
	<meta property="og:image:width" content="948">
	<meta property="og:image:height" content="424">
	<meta property="og:type" content="website" />

	<link rel="stylesheet" href="css/reset.css">
    <link rel="stylesheet" href="css/style.css?<?php echo rand(0,1223131) ?>">

    <title><?php echo $meta_title; ?></title>

</head>
<body>
	<div id="disco">
		<?php
			$myRand = rand(1,2);
			
			if($myRand == 1){
				?>
				<a href="https://open.spotify.com/album/6WMQfW07qWUaFJpp298uPi" target="_blank"><img src="images/Capa_TellMeYourName.jpg?<?php //echo rand(0,1223131) ?>" /><br />
		Tell Me Your Name</a><br />
				<?php
			}

			
			if($myRand == 2){
				?>
				<a href="https://open.spotify.com/album/1o2nq3TbCGmFmgsBth1XL1" target="_blank"><img src="images/Capa_StreetProsecution.jpg?<?php //echo rand(0,1223131) ?>" /><br />
		Street Prosecution</a><br />
				<?php
			}
			
		?>
		

	</div>
	<a href="https://open.spotify.com/playlist/0PmQC52w1gm5TY9le9IyLt" target="_blank"><img src="images/logo.png?<?php echo rand(0,1223131) ?>" alt="Cesar Bouli" class="logo" /></a>
</body>
<script>
window.location.replace("https://open.spotify.com/playlist/280Kv8R7Js4QLIF8o94YlZ");
</script>
</html>