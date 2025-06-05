<?php

use App\Http\Controllers\UserController;
use Illuminate\Support\Facades\Route;

Route::group(
	[
		'prefix' => 'users',
		'middleware' => 'auth:sanctum',
	],
	function()
	{
		Route::put('/me', [UserController::class, 'update']);
	}
);