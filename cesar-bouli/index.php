<?php
$meta_image = "http://cesarbouli.com/images/cesar-bouli/photo.jpg";
$meta_text = "Dia 12 de maio, lanço o meu primeiro disco solo entitulado 'Cesar Bouli'! Dá aquela força nas redes sociais abaixo e faça seu pre-save!";
$meta_title = "Cesar Bouli - Hey, Irmão";
$root_url = "http://cesarbouli.com";
?><!DOCTYPE html>
<html lang="es">
<head>
	<script>
	window.location.replace("https://distrokid.com/hyperfollow/cesarbouli/cesar-bouli");
	</script>
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

	<meta name="msapplication-TileColor" content="#000000">

	<meta name="theme-color" content="#000000">
	<meta name="description" content="<?php echo $meta_text; ?>" />

	<meta property="og:locale" content="pt-br" />
	<meta property="og:url" content="http://cesarbouli.com/hey-irmao">
	<meta property="og:title" content="<?php echo $meta_title; ?>">
	<meta property="og:site_name" content="<?php echo $meta_title; ?>">
	<meta property="og:description" content="<?php echo $meta_text; ?>">
	<meta property="og:image" content="<?php echo $meta_image; ?>">
	<meta property="og:image:type" content="image/png">
	<meta property="og:image:width" content="948">
	<meta property="og:image:height" content="424">
	<meta property="og:type" content="website" />

	<link rel="stylesheet" href="<?php echo $root_url; ?>css/reset.css">
    <link rel="stylesheet" href="<?php echo $root_url; ?>css/style.css?<?php echo rand(0,1223131) ?>">

    <title><?php echo $meta_title; ?></title>

</head>
<body>
	<div id="disco">
		<?php
			/*$myRand = rand(1,2);

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
			*/
		?>


	</div>
	<a href="https://open.spotify.com/playlist/0PmQC52w1gm5TY9le9IyLt" target="_blank"><img src="images/logo.png?<?php echo rand(0,1223131) ?>" alt="Cesar Bouli" class="logo" /></a>
</body>
</html>
