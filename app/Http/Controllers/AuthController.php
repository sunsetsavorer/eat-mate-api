<?php

namespace App\Http\Controllers;

use App\Domains\Auth\UseCases\SigninUseCase;
use App\Http\Requests\Auth\SigninRequest;
use App\Http\Resources\Auth\SigninResource;
use App\Repositories\Auth\AuthTokenRepository;
use App\Repositories\User\UserRepository;

class AuthController extends Controller
{
	public function signin(SigninRequest $request): SigninResource
	{
		$data = $request->validated();

		$uc = new SigninUseCase(
			new AuthTokenRepository(),
			new UserRepository()
		);

		$response = $uc(
			$data['telegramID']
		);

		return new SigninResource($response);
	}
}
