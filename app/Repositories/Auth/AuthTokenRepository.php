<?php

namespace App\Repositories\Auth;

use App\Domains\Auth\Interfaces\AuthTokenRepositoryInterface;
use App\Domains\Common\Exceptions\NotFoundException;
use App\Models\User;

class AuthTokenRepository implements AuthTokenRepositoryInterface
{
	public function createByUserID(string $userID): string
	{
		$user = User::find($userID);

		if (!$user)
			throw new NotFoundException();

		$token = $user->createToken('AuthToken')->plainTextToken;

		return $token;
	}
}