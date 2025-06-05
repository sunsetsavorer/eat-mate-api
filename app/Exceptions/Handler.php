<?php

namespace App\Exceptions;

use App\Domains\Common\Exceptions\ServiceException;
use Illuminate\Foundation\Exceptions\Handler as ExceptionHandler;
use Illuminate\Validation\ValidationException;
use Throwable;

class Handler extends ExceptionHandler
{
	/**
	 * The list of the inputs that are never flashed to the session on validation exceptions.
	 *
	 * @var array<int, string>
	 */
	protected $dontFlash = [
		'current_password',
		'password',
		'password_confirmation',
	];

	public function register(): void
	{
		$this->reportable(function (Throwable $e) {
			//
		});
	}

	public function render($request, Throwable $exception)
	{
		list($message, $code) = match (get_class($exception)) {
			ServiceException::class => [
				['other' => $exception->getMessage()],
				$exception->getCode(),
			],
			ValidationException::class => [
				$exception->validator->errors()->getMessages(),
				422,
			],
			default => [
				$exception->getMessage(),
				400,
			],
		};

		return response()->json([
			'errors' => $message,
		], $code);
	}
}
