<?php

namespace App\Domains\User\Interfaces;

use App\Domains\User\Entities\UserEntity;

interface UserRepositoryInterface
{
	public function create(UserEntity $entity): string;
	public function getByTelegramID(int $telegramID): UserEntity;
}