<?php

namespace App\Domains\User\UseCases;

use App\Domains\Common\Exceptions\NotFoundException;
use App\Domains\Common\Exceptions\ServiceException;
use App\Domains\User\DTOs\UpdateUserDTO;
use App\Domains\User\Entities\UserEntity;
use App\Domains\User\Interfaces\UserRepositoryInterface;

class UpdateUserUseCase
{
	public function __construct(
		private UserRepositoryInterface $userRepository,
	){}

	public function __invoke(UpdateUserDTO $dto): void
	{
		try {
			$user = new UserEntity();

			$user->setTelegramID($dto->getTelegramID());
			$user->setName($dto->getName());
			$user->setPhotoUrl($dto->getPhotoUrl());

			$this->userRepository->updateByTelegramID($user);

		} catch (NotFoundException $e) {
			throw new ServiceException('Пользователь не найден', 404);
		}
	}
}