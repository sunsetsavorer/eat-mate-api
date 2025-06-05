<?php

namespace App\Repositories\User;

use App\Domains\Common\Exceptions\NotFoundException;
use App\Domains\User\Entities\UserEntity;
use App\Domains\User\Interfaces\UserRepositoryInterface;
use App\Models\User;

class UserRepository implements UserRepositoryInterface
{
	public function create(UserEntity $entity): string
	{
		$user = User::create([
			'telegram_id' => $entity->getTelegramID(),
			'name' => $entity->getName(),
			'photo_url' => $entity->getPhotoUrl(),
		]);

		return $user->id;
	}

	public function getByTelegramID(int $telegramID): UserEntity
	{
		$user = User::query()
			->where('telegram_id', $telegramID)
			->first();

		if (!$user)
			throw new NotFoundException();

		$userEntity = new UserEntity();

		$userEntity->setID($user->id);
		$userEntity->setTelegramID($user->telegram_id);
		$userEntity->setName($user->name);
		$userEntity->setPhotoUrl($user->photo_url);

		return $userEntity;
	}

	public function updateByTelegramID(UserEntity $entity): void
	{
		$user = User::query()
			->where('telegram_id', $entity->getTelegramID())
			->first();

		if (!$user)
			throw new NotFoundException();

		$user->update([
			'name' => $entity->getName(),
			'photo_url' => $entity->getPhotoUrl(),
		]);
	}
}