<?php
set_time_limit(0);

$host = "127.0.0.1";
$port = 8080;
$socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP)or die("Could not create  socket\n");

$connection = socket_connect($socket, $host, $port) or die("Could not connet server\n");
socket_write($socket, "*2\r\n$3\r\nGET\r\n$5\r\nHENRY\r\n") or die("Write failed\n");

socket_close($socket);
