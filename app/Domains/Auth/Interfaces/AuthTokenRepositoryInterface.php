<?php

namespace App\Domains\Auth\Interfaces;

interface AuthTokenRepositoryInterface
{
	public function createByUserID(string $userID): string;
}