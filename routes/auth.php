<?php

use App\Http\Controllers\AuthController;
use Illuminate\Support\Facades\Route;

Route::group(
	[
		'prefix' => 'auth'
	],
	function()
	{
		Route::post('/signin', [AuthController::class, 'signin']);
	}
);