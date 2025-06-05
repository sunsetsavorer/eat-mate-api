<?php

namespace App\Http\Controllers;

use App\Domains\User\DTOs\UpdateUserDTO;
use App\Domains\User\UseCases\UpdateUserUseCase;
use App\Http\Requests\User\UpdateUserRequest;
use App\Http\Resources\User\UpdateUserResource;
use App\Repositories\User\UserRepository;

class UserController extends Controller
{
    public function update(UpdateUserRequest $request)
	{
		$data = $request->validated();
		$user = request()->user();

		$uc = new UpdateUserUseCase(
			new UserRepository()
		);

		$uc(
			new UpdateUserDTO(
				$user->telegram_id,
				$data['name'],
				$data['photoUrl']
			)
		);

		return new UpdateUserResource([]);
	}
}
