<?php

namespace App\Domains\Auth\UseCases;

use App\Domains\Auth\Interfaces\AuthTokenRepositoryInterface;
use App\Domains\Common\Exceptions\NotFoundException;
use App\Domains\User\Entities\UserEntity;
use App\Domains\User\Interfaces\UserRepositoryInterface;

class SigninUseCase
{
	public function __construct(
		private AuthTokenRepositoryInterface $authTokenRepository,
		private UserRepositoryInterface $userRepository,
	){}

	public function __invoke(int $telegramID): string
	{
		try {
			$user = $this->userRepository->getByTelegramID($telegramID);
			$token = $this->authTokenRepository->createByUserID($user->getID());

		} catch (NotFoundException $e) {
			$user = new UserEntity();

			$user->setTelegramID($telegramID);

			$userID = $this->userRepository->create($user);

			$token = $this->authTokenRepository->createByUserID($userID);
		}

		return $token;
	}
}